// Code generated by hertz generator. DO NOT EDIT.

package core_api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	core_api "github.com/xh-polaris/meowchat-core-api/biz/adaptor/controller/core_api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/get_min_version", append(_getminversionMw(), core_api.GetMinVersion)...)
	root.GET("/prefetch", append(_prefetchMw(), core_api.Prefetch)...)
	{
		_auth := root.Group("/auth", _authMw()...)
		_auth.POST("/send_verify_code", append(_sendverifycodeMw(), core_api.SendVerifyCode)...)
		_auth.POST("/set_password", append(_setpasswordMw(), core_api.SetPassword)...)
		_auth.POST("/sign_in", append(_signinMw(), core_api.SignIn)...)
	}
	{
		_collection := root.Group("/collection", _collectionMw()...)
		_collection.POST("/create_image", append(_createimageMw(), core_api.CreateImage)...)
		_collection.POST("/delete_cat", append(_deletecatMw(), core_api.DeleteCat)...)
		_collection.POST("/delete_image", append(_deleteimageMw(), core_api.DeleteImage)...)
		_collection.GET("/get_cat_detail", append(_getcatdetailMw(), core_api.GetCatDetail)...)
		_collection.GET("/get_cat_previews", append(_getcatpreviewsMw(), core_api.GetCatPreviews)...)
		_collection.GET("/get_image_by_cat", append(_getimagebycatMw(), core_api.GetImageByCat)...)
		_collection.POST("/new_cat", append(_newcatMw(), core_api.NewCat)...)
		_collection.GET("/search_cat", append(_searchcatMw(), core_api.SearchCat)...)
	}
	{
		_comment := root.Group("/comment", _commentMw()...)
		_comment.POST("/delete_comment", append(_deletecommentMw(), core_api.DeleteComment)...)
		_comment.GET("/get_comments", append(_getcommentsMw(), core_api.GetComments)...)
		_comment.POST("/new_comment", append(_newcommentMw(), core_api.NewComment)...)
	}
	{
		_community := root.Group("/community", _communityMw()...)
		_community.POST("/delete_community", append(_deletecommunityMw(), core_api.DeleteCommunity)...)
		_community.GET("/list_community", append(_listcommunityMw(), core_api.ListCommunity)...)
		_community.POST("/new_community", append(_newcommunityMw(), core_api.NewCommunity)...)
	}
	{
		_like := root.Group("/like", _likeMw()...)
		_like.POST("/do_like", append(_dolikeMw(), core_api.DoLike)...)
		_like.GET("/get_count", append(_getlikedcountMw(), core_api.GetLikedCount)...)
		_like.GET("/get_liked_users", append(_getlikedusersMw(), core_api.GetLikedUsers)...)
		_like.GET("/get_user_like_contents", append(_getuserlikecontentsMw(), core_api.GetUserLikeContents)...)
		_like.GET("/get_user_liked", append(_getuserlikedMw(), core_api.GetUserLiked)...)
		_like.GET("/get_user_likes", append(_getuserlikesMw(), core_api.GetUserLikes)...)
	}
	{
		_moment := root.Group("/moment", _momentMw()...)
		_moment.POST("/delete_moment", append(_deletemomentMw(), core_api.DeleteMoment)...)
		_moment.GET("/get_moment_detail", append(_getmomentdetailMw(), core_api.GetMomentDetail)...)
		_moment.GET("/get_moment_previews", append(_getmomentpreviewsMw(), core_api.GetMomentPreviews)...)
		_moment.POST("/new_moment", append(_newmomentMw(), core_api.NewMoment)...)
		_moment.GET("/search_moment", append(_searchmomentMw(), core_api.SearchMoment)...)
	}
	{
		_notice := root.Group("/notice", _noticeMw()...)
		_notice.POST("/delete_admin", append(_deleteadminMw(), core_api.DeleteAdmin)...)
		_notice.GET("/get_admins", append(_getadminsMw(), core_api.GetAdmins)...)
		_notice.GET("/get_news", append(_getnewsMw(), core_api.GetNews)...)
		_notice.GET("/get_notices", append(_getnoticesMw(), core_api.GetNotices)...)
		_notice.POST("/new_admin", append(_newadminMw(), core_api.NewAdmin)...)
		_notice.POST("/new_news", append(_newnewsMw(), core_api.NewNews)...)
		_notice.POST("/new_notice", append(_newnoticeMw(), core_api.NewNotice)...)
		_notice.POST("/remove_news", append(_deletenewsMw(), core_api.DeleteNews)...)
		_notice.POST("/remove_notice", append(_deletenoticeMw(), core_api.DeleteNotice)...)
	}
	{
		_notification := root.Group("/notification", _notificationMw()...)
		_notification.GET("/clean_notification", append(_cleannotificationMw(), core_api.CleanNotification)...)
		_notification.GET("/count_notification", append(_countnotificationMw(), core_api.CountNotification)...)
		_notification.GET("/list_notification", append(_listnotificationMw(), core_api.ListNotification)...)
		_notification.GET("/read_notification", append(_readnotificationMw(), core_api.ReadNotification)...)
	}
	{
		_plan := root.Group("/plan", _planMw()...)
		_plan.GET("/count_donate_by_plan", append(_countdonatebyplanMw(), core_api.CountDonateByPlan)...)
		_plan.GET("/count_donate_by_user", append(_countdonatebyuserMw(), core_api.CountDonateByUser)...)
		_plan.POST("/delete_plan", append(_deleteplanMw(), core_api.DeletePlan)...)
		_plan.GET("/donate_fish", append(_donatefishMw(), core_api.DonateFish)...)
		_plan.GET("/get_plan_detail", append(_getplandetailMw(), core_api.GetPlanDetail)...)
		_plan.GET("/get_plan_previews", append(_getplanpreviewsMw(), core_api.GetPlanPreviews)...)
		_plan.GET("/get_user_fish", append(_getuserfishMw(), core_api.GetUserFish)...)
		_plan.GET("/list_donate_by_user", append(_listdonatebyuserMw(), core_api.ListDonateByUser)...)
		_plan.GET("/list_fish_by_plan", append(_listfishbyplanMw(), core_api.ListFishByPlan)...)
		_plan.POST("/new_plan", append(_newplanMw(), core_api.NewPlan)...)
	}
	{
		_post := root.Group("/post", _postMw()...)
		_post.POST("/delete_post", append(_deletepostMw(), core_api.DeletePost)...)
		_post.GET("/get_post_detail", append(_getpostdetailMw(), core_api.GetPostDetail)...)
		_post.POST("/get_post_previews", append(_getpostpreviewsMw(), core_api.GetPostPreviews)...)
		_post.POST("/new_post", append(_newpostMw(), core_api.NewPost)...)
		_post.POST("/set_official", append(_setofficialMw(), core_api.SetOfficial)...)
	}
	{
		_role := root.Group("/role", _roleMw()...)
		_role.POST("/create_apply", append(_createapplyMw(), core_api.CreateApply)...)
		_role.GET("/get_user_by_role", append(_getuserbyroleMw(), core_api.GetUserByRole)...)
		_role.GET("/get_user_roles", append(_getuserrolesMw(), core_api.GetUserRoles)...)
		_role.POST("/handle_apply", append(_handleapplyMw(), core_api.HandleApply)...)
		_role.GET("/list_apply", append(_listapplyMw(), core_api.ListApply)...)
		_role.POST("/update_community_admin", append(_updatecommunityadminMw(), core_api.UpdateCommunityAdmin)...)
		_role.POST("/update_role", append(_updateroleMw(), core_api.UpdateRole)...)
		_role.POST("/update_super_admin", append(_updatesuperadminMw(), core_api.UpdateSuperAdmin)...)
	}
	{
		_sts := root.Group("/sts", _stsMw()...)
		_sts.POST("/apply_signed_url", append(_applysignedurlMw(), core_api.ApplySignedUrl)...)
		_sts.POST("/apply_signed_url_as_community", append(_applysignedurlascommunityMw(), core_api.ApplySignedUrlAsCommunity)...)
	}
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/check_in", append(_checkinMw(), core_api.CheckIn)...)
		_user.GET("/get_user_info", append(_getuserinfoMw(), core_api.GetUserInfo)...)
		_user.GET("/search_user", append(_searchuserMw(), core_api.SearchUser)...)
		_user.POST("/update_user_info", append(_updateuserinfoMw(), core_api.UpdateUserInfo)...)
	}
}
