package service

import (
	"context"

	"github.com/google/wire"
	"github.com/samber/lo"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/basic"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/content"
	genlike "github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowchat/user"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/comment"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/core_api"
	"github.com/xh-polaris/meowchat-core-api/biz/application/dto/meowchat/user"
	"github.com/xh-polaris/meowchat-core-api/biz/domain/service"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/config"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_content"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/meowchat_user"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/rpc/platform_comment"
	"github.com/xh-polaris/meowchat-core-api/biz/infrastructure/util"
)

type ILikeService interface {
	DoLike(ctx context.Context, req *core_api.DoLikeReq, user *basic.UserMeta) (*core_api.DoLikeResp, error)
	GetLikedCount(ctx context.Context, req *core_api.GetLikedCountReq) (*core_api.GetLikedCountResp, error)
	GetLikedUsers(ctx context.Context, req *core_api.GetLikedUsersReq) (*core_api.GetLikedUsersResp, error)
	GetUserLiked(ctx context.Context, req *core_api.GetUserLikedReq, user *basic.UserMeta) (*core_api.GetUserLikedResp, error)
	GetUserLikes(ctx context.Context, req *core_api.GetUserLikesReq) (*core_api.GetUserLikesResp, error)
	GetUserLikeContents(ctx context.Context, req *core_api.GetUserLikeContentsReq) (*core_api.GetUserLikeContentsResp, error)
}

type LikeService struct {
	Config               *config.Config
	User                 meowchat_user.IMeowchatUser
	Content              meowchat_content.IMeowchatContent
	Comment              platform_comment.IPlatformCommment
	PostDomainService    service.IPostDomainService
	MomentDomainService  service.IMomentDomainService
	CommentDomainService service.ICommentDomainService
}

var LikeServiceSet = wire.NewSet(
	wire.Struct(new(LikeService), "*"),
	wire.Bind(new(ILikeService), new(*LikeService)),
)

func (s *LikeService) DoLike(ctx context.Context, req *core_api.DoLikeReq, user *basic.UserMeta) (*core_api.DoLikeResp, error) {
	resp := new(core_api.DoLikeResp)

	userId := user.UserId
	associatedId := ""
	likedUserId := ""

	if req.GetTargetType() == 1 {
		data, err := s.Content.RetrievePost(ctx, &content.RetrievePostReq{PostId: req.TargetId})
		if err != nil {
			return nil, err
		}
		associatedId = data.Post.Id
		likedUserId = data.Post.UserId
	} else if req.GetTargetType() == 4 {
		data, err := s.Content.RetrieveMoment(ctx, &content.RetrieveMomentReq{MomentId: req.TargetId})
		if err != nil {
			return nil, err
		}
		associatedId = data.Moment.Id
		likedUserId = data.Moment.UserId
	} else if req.GetTargetType() == 2 {
		data, err := s.Comment.RetrieveCommentById(ctx, &comment.RetrieveCommentByIdReq{Id: req.TargetId})
		if err != nil {
			return nil, err
		}
		associatedId = data.Comment.ParentId
		likedUserId = data.Comment.AuthorId
	} else if req.GetTargetType() == 6 {
		associatedId = req.TargetId
		likedUserId = req.TargetId
	}

	r, err := s.User.DoLike(ctx, &genlike.DoLikeReq{
		UserId:       userId,
		TargetId:     req.TargetId,
		Type:         genlike.LikeType(req.TargetType),
		AssociatedId: associatedId,
		LikedUserId:  likedUserId,
	})

	if err != nil {
		return nil, err
	}
	if r.GetGetFish() == true {
		_, err = s.Content.AddUserFish(ctx, &content.AddUserFishReq{
			UserId: user.UserId,
			Fish:   s.Config.Fish.Like[r.GetFishTimes-1],
		})
		if err == nil {
			resp.GetFishNum = s.Config.Fish.Like[r.GetFishTimes-1]
		}
	}
	resp.GetFish = r.GetGetFish()
	return resp, nil
}

func (s *LikeService) GetLikedCount(ctx context.Context, req *core_api.GetLikedCountReq) (*core_api.GetLikedCountResp, error) {
	resp := new(core_api.GetLikedCountResp)

	likes, err := s.User.GetTargetLikes(ctx, &genlike.GetTargetLikesReq{
		TargetId: req.TargetId,
		Type:     genlike.LikeType(req.TargetType),
	})
	if err != nil {
		return nil, err
	}

	resp.Count = likes.Count

	return resp, nil
}

func (s *LikeService) GetLikedUsers(ctx context.Context, req *core_api.GetLikedUsersReq) (*core_api.GetLikedUsersResp, error) {
	resp := new(core_api.GetLikedUsersResp)
	data, err := s.User.GetLikedUsers(ctx, &genlike.GetLikedUsersReq{
		TargetId: req.TargetId,
		Type:     genlike.LikeType(req.TargetType),
	})
	if err != nil {
		return nil, err
	}
	resp.Users = make([]*user.UserPreview, 0, len(data.UserIds))
	for _, userId := range data.UserIds {
		res, err := s.User.GetUser(ctx, &genlike.GetUserReq{UserId: userId})
		if err != nil {
			logx.Error(err)
		}
		resp.Users = append(resp.Users, &user.UserPreview{
			Id:        res.User.Id,
			Nickname:  res.User.Nickname,
			AvatarUrl: res.User.AvatarUrl,
		})
	}
	return resp, nil
}

func (s *LikeService) GetUserLiked(ctx context.Context, req *core_api.GetUserLikedReq, user *basic.UserMeta) (*core_api.GetUserLikedResp, error) {
	resp := new(core_api.GetUserLikedResp)

	userId := user.UserId
	like, err := s.User.GetUserLike(ctx, &genlike.GetUserLikedReq{
		UserId:   userId,
		TargetId: req.TargetId,
		Type:     genlike.LikeType(req.TargetType),
	})
	if err != nil {
		return nil, err
	}

	resp.Liked = like.Liked

	return resp, nil
}

func (s *LikeService) GetUserLikes(ctx context.Context, req *core_api.GetUserLikesReq) (*core_api.GetUserLikesResp, error) {
	resp := new(core_api.GetUserLikesResp)
	data, err := s.User.GetUserLikes(ctx, &genlike.GetUserLikesReq{
		UserId: req.UserId,
		Type:   genlike.LikeType(req.TargetType),
	})
	if err != nil {
		return nil, err
	}
	resp.Likes = make([]*user.Like, 0, len(data.Likes))
	for _, like := range data.Likes {
		resp.Likes = append(resp.Likes, &user.Like{
			TargetId:     like.TargetId,
			AssociatedId: like.AssociatedId,
		})
	}
	return resp, nil
}

func (s *LikeService) GetUserLikeContents(ctx context.Context, req *core_api.GetUserLikeContentsReq) (*core_api.GetUserLikeContentsResp, error) {
	resp := new(core_api.GetUserLikeContentsResp)
	data, err := s.User.GetUserLikes(ctx, &genlike.GetUserLikesReq{
		UserId: req.UserId,
		Type:   genlike.LikeType(req.TargetType),
	})
	if err != nil {
		return nil, err
	}
	if req.GetTargetType() == 1 { //post
		resp.Posts = make([]*core_api.Post, len(data.Likes))
		util.ParallelRun(lo.Map(data.Likes, func(like *genlike.Like, i int) func() {
			return func() {
				post, err := s.Content.RetrievePost(ctx, &content.RetrievePostReq{PostId: like.TargetId})
				if err == nil {
					p := &core_api.Post{
						Id:         post.Post.Id,
						CreateAt:   post.Post.CreateAt,
						Title:      post.Post.Title,
						Text:       post.Post.Text,
						CoverUrl:   lo.EmptyableToPtr(post.Post.CoverUrl),
						Tags:       post.Post.Tags,
						IsOfficial: post.Post.IsOfficial,
					}
					util.ParallelRun([]func(){
						func() {
							_ = s.PostDomainService.LoadAuthor(ctx, p, post.Post.UserId)
						},
						func() {
							_ = s.PostDomainService.LoadLikeCount(ctx, p)
						},
						func() {
							_ = s.PostDomainService.LoadCommentCount(ctx, p)
						},
					})
					resp.Posts[i] = p
				}
			}
		}))
	} else if req.GetTargetType() == 4 { //moment
		resp.Moments = make([]*core_api.Moment, len(data.Likes))
		util.ParallelRun(lo.Map(data.Likes, func(like *genlike.Like, i int) func() {
			return func() {
				moment, err := s.Content.RetrieveMoment(ctx, &content.RetrieveMomentReq{MomentId: like.TargetId})
				if err == nil {
					resp.Moments[i] = &core_api.Moment{
						Id:          moment.Moment.Id,
						CreateAt:    moment.Moment.CreateAt,
						Photos:      moment.Moment.Photos,
						Title:       moment.Moment.Title,
						Text:        moment.Moment.Text,
						CommunityId: moment.Moment.CommunityId,
					}
					if err == nil {
						util.ParallelRun([]func(){
							func() {
								if moment.Moment.GetCatId() != "" {
									_ = s.MomentDomainService.LoadCats(ctx, resp.Moments[i], []string{moment.Moment.GetCatId()})
								}
							},
							func() {
								_ = s.MomentDomainService.LoadAuthor(ctx, resp.Moments[i], moment.Moment.UserId)
							},
							func() {
								_ = s.MomentDomainService.LoadLikeCount(ctx, resp.Moments[i])
							},
							func() {
								_ = s.MomentDomainService.LoadCommentCount(ctx, resp.Moments[i])
							},
						})
					}
				}
			}
		}))
	} else if req.GetTargetType() == 2 { //comment
		resp.Comments = make([]*core_api.Comment, len(data.Likes))
		util.ParallelRun(lo.Map(data.Likes, func(like *genlike.Like, i int) func() {
			return func() {
				_comment, err := s.Comment.RetrieveCommentById(ctx, &comment.RetrieveCommentByIdReq{Id: like.TargetId})
				if err == nil {
					c := &core_api.Comment{
						Id:       _comment.Comment.Id,
						CreateAt: _comment.Comment.CreateAt,
						Text:     _comment.Comment.Text,
					}
					util.ParallelRun([]func(){
						func() {
							if _comment.Comment.ReplyTo == "" {
								return
							}
							_ = s.CommentDomainService.LoadReplyUser(ctx, c, _comment.Comment.ReplyTo)
						},
						func() {
							_ = s.CommentDomainService.LoadCommentCount(ctx, c)
						},
						func() {
							_ = s.CommentDomainService.LoadLikeCount(ctx, c)
						},
						func() {
							_ = s.CommentDomainService.LoadAuthor(ctx, c, _comment.Comment.AuthorId)
						},
					})
					resp.Comments[i] = c
				}
			}
		}))
	} else if req.GetTargetType() == 6 { //user
		resp.Users = make([]*user.UserPreview, len(data.Likes))
		util.ParallelRun(lo.Map(data.Likes, func(like *genlike.Like, i int) func() {
			return func() {
				_user, err := s.User.GetUserDetail(ctx, &genlike.GetUserDetailReq{UserId: like.TargetId})
				if err == nil {
					resp.Users[i] = &user.UserPreview{
						Id:        _user.User.Id,
						Nickname:  _user.User.Nickname,
						AvatarUrl: _user.User.AvatarUrl,
					}
				}
			}
		}))
	}
	return resp, nil
}
