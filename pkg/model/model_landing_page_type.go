/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// LandingPageType : 落地页页面类型
type LandingPageType string

// List of LandingPageType
const (
	LandingPageType_LANDING_PAGE_TYPE_UNKNOWN                         LandingPageType = "LANDING_PAGE_TYPE_UNKNOWN"
	LandingPageType_LANDING_PAGE_TYPE_ANDROID_APP                     LandingPageType = "LANDING_PAGE_TYPE_ANDROID_APP"
	LandingPageType_LANDING_PAGE_TYPE_IOS_APP                         LandingPageType = "LANDING_PAGE_TYPE_IOS_APP"
	LandingPageType_LANDING_PAGE_TYPE_FENGYE_WEB                      LandingPageType = "LANDING_PAGE_TYPE_FENGYE_WEB"
	LandingPageType_LANDING_PAGE_TYPE_FENGYE_LBS                      LandingPageType = "LANDING_PAGE_TYPE_FENGYE_LBS"
	LandingPageType_LANDING_PAGE_TYPE_NATIVE                          LandingPageType = "LANDING_PAGE_TYPE_NATIVE"
	LandingPageType_LANDING_PAGE_TYPE_ANDROID_APP_NATIVE              LandingPageType = "LANDING_PAGE_TYPE_ANDROID_APP_NATIVE"
	LandingPageType_LANDING_PAGE_TYPE_IOS_APP_NATIVE                  LandingPageType = "LANDING_PAGE_TYPE_IOS_APP_NATIVE"
	LandingPageType_LANDING_PAGE_TYPE_XJ_ANDROID_APP_H5               LandingPageType = "LANDING_PAGE_TYPE_XJ_ANDROID_APP_H5"
	LandingPageType_LANDING_PAGE_TYPE_XJ_IOS_APP_H5                   LandingPageType = "LANDING_PAGE_TYPE_XJ_IOS_APP_H5"
	LandingPageType_LANDING_PAGE_TYPE_XJ_WEBSITE_H5                   LandingPageType = "LANDING_PAGE_TYPE_XJ_WEBSITE_H5"
	LandingPageType_LANDING_PAGE_TYPE_XJ_WEBSITE_NATIVE               LandingPageType = "LANDING_PAGE_TYPE_XJ_WEBSITE_NATIVE"
	LandingPageType_LANDING_PAGE_TYPE_XJ_ANDROID_APP_NATIVE           LandingPageType = "LANDING_PAGE_TYPE_XJ_ANDROID_APP_NATIVE"
	LandingPageType_LANDING_PAGE_TYPE_XJ_IOS_APP_NATIVE               LandingPageType = "LANDING_PAGE_TYPE_XJ_IOS_APP_NATIVE"
	LandingPageType_LANDING_PAGE_TYPE_FENGYE_EC_SINGLE                LandingPageType = "LANDING_PAGE_TYPE_FENGYE_EC_SINGLE"
	LandingPageType_LANDING_PAGE_TYPE_FENGYE_EC_TOGETHER              LandingPageType = "LANDING_PAGE_TYPE_FENGYE_EC_TOGETHER"
	LandingPageType_LANDING_PAGE_TYPE_FENGYE_EC_FOCUS                 LandingPageType = "LANDING_PAGE_TYPE_FENGYE_EC_FOCUS"
	LandingPageType_LANDING_PAGE_TYPE_FENGYE_EC_COMMON                LandingPageType = "LANDING_PAGE_TYPE_FENGYE_EC_COMMON"
	LandingPageType_LANDING_PAGE_TYPE_FENGYE_EC_WECHAT_MINIPROGRAM    LandingPageType = "LANDING_PAGE_TYPE_FENGYE_EC_WECHAT_MINIPROGRAM"
	LandingPageType_LANDING_PAGE_TYPE_PLAY_H5                         LandingPageType = "LANDING_PAGE_TYPE_PLAY_H5"
	LandingPageType_LANDING_PAGE_TYPE_YOUZAN_SINGLE                   LandingPageType = "LANDING_PAGE_TYPE_YOUZAN_SINGLE"
	LandingPageType_LANDING_PAGE_TYPE_YOUZAN_TOGETHER                 LandingPageType = "LANDING_PAGE_TYPE_YOUZAN_TOGETHER"
	LandingPageType_LANDING_PAGE_TYPE_YOUZAN_WECHAT_MINIPROGRAM       LandingPageType = "LANDING_PAGE_TYPE_YOUZAN_WECHAT_MINIPROGRAM"
	LandingPageType_LANDING_PAGE_TYPE_YIYE_FORM                       LandingPageType = "LANDING_PAGE_TYPE_YIYE_FORM"
	LandingPageType_LANDING_PAGE_TYPE_JINSHUJU_FORM                   LandingPageType = "LANDING_PAGE_TYPE_JINSHUJU_FORM"
	LandingPageType_LANDING_PAGE_TYPE_WEIMOB_PRODUCTSET               LandingPageType = "LANDING_PAGE_TYPE_WEIMOB_PRODUCTSET"
	LandingPageType_LANDING_PAGE_TYPE_WEIMOB_PROMOTION                LandingPageType = "LANDING_PAGE_TYPE_WEIMOB_PROMOTION"
	LandingPageType_LANDING_PAGE_TYPE_WEIMOB_PRODUCT                  LandingPageType = "LANDING_PAGE_TYPE_WEIMOB_PRODUCT"
	LandingPageType_LANDING_PAGE_TYPE_WEIMOB_H5                       LandingPageType = "LANDING_PAGE_TYPE_WEIMOB_H5"
	LandingPageType_LANDING_PAGE_TYPE_WEIMOB_WECHAT_MINIPROGRAM       LandingPageType = "LANDING_PAGE_TYPE_WEIMOB_WECHAT_MINIPROGRAM"
	LandingPageType_LANDING_PAGE_TYPE_QQ_MOBILE_MINI_PROGRAM          LandingPageType = "LANDING_PAGE_TYPE_QQ_MOBILE_MINI_PROGRAM"
	LandingPageType_LANDING_PAGE_TYPE_QQ_BROWSER_MINI_PROGRAM         LandingPageType = "LANDING_PAGE_TYPE_QQ_BROWSER_MINI_PROGRAM"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_NATIVE                   LandingPageType = "LANDING_PAGE_TYPE_WECHAT_NATIVE"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_H5                       LandingPageType = "LANDING_PAGE_TYPE_WECHAT_H5"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_OFFICIAL_ACCOUNT_ARTICLE LandingPageType = "LANDING_PAGE_TYPE_WECHAT_OFFICIAL_ACCOUNT_ARTICLE"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_COUPON                   LandingPageType = "LANDING_PAGE_TYPE_WECHAT_COUPON"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_OFFICIAL_ACCOUNT_DETAIL  LandingPageType = "LANDING_PAGE_TYPE_WECHAT_OFFICIAL_ACCOUNT_DETAIL"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_APPSTORE_MOMENTS         LandingPageType = "LANDING_PAGE_TYPE_WECHAT_APPSTORE_MOMENTS"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_H5_TEMPLATE              LandingPageType = "LANDING_PAGE_TYPE_WECHAT_H5_TEMPLATE"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_MOMENTS_SIMPLE_NATIVE    LandingPageType = "LANDING_PAGE_TYPE_WECHAT_MOMENTS_SIMPLE_NATIVE"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_FULL_SCREEN              LandingPageType = "LANDING_PAGE_TYPE_WECHAT_FULL_SCREEN"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_BASE_DETAILS             LandingPageType = "LANDING_PAGE_TYPE_WECHAT_BASE_DETAILS"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_PICTURE_DETAILS          LandingPageType = "LANDING_PAGE_TYPE_WECHAT_PICTURE_DETAILS"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_STORE                    LandingPageType = "LANDING_PAGE_TYPE_WECHAT_STORE"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_MOMENTS_H5_BRAND         LandingPageType = "LANDING_PAGE_TYPE_WECHAT_MOMENTS_H5_BRAND"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_MOMENTS_PICTURE_FORM     LandingPageType = "LANDING_PAGE_TYPE_WECHAT_MOMENTS_PICTURE_FORM"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_MOMENTS_VIDEO_FORM       LandingPageType = "LANDING_PAGE_TYPE_WECHAT_MOMENTS_VIDEO_FORM"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_MOMENTS_APPLIED_TEXT     LandingPageType = "LANDING_PAGE_TYPE_WECHAT_MOMENTS_APPLIED_TEXT"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_MOMENTS_APPLIED_VIDEO    LandingPageType = "LANDING_PAGE_TYPE_WECHAT_MOMENTS_APPLIED_VIDEO"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_FORM_TEMPLATE            LandingPageType = "LANDING_PAGE_TYPE_WECHAT_FORM_TEMPLATE"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_VOTING_TEMPLATE          LandingPageType = "LANDING_PAGE_TYPE_WECHAT_VOTING_TEMPLATE"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_PAYMENT_COUPONS          LandingPageType = "LANDING_PAGE_TYPE_WECHAT_PAYMENT_COUPONS"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_MINI_GAME                LandingPageType = "LANDING_PAGE_TYPE_WECHAT_MINI_GAME"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_NATIVE_BETA              LandingPageType = "LANDING_PAGE_TYPE_WECHAT_NATIVE_BETA"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_CALL_DAILOG              LandingPageType = "LANDING_PAGE_TYPE_WECHAT_CALL_DAILOG"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_FOCUS_DAILOG             LandingPageType = "LANDING_PAGE_TYPE_WECHAT_FOCUS_DAILOG"
	LandingPageType_LANDING_PAGE_TYPE_WECHAT_MINI_PROGRAM             LandingPageType = "LANDING_PAGE_TYPE_WECHAT_MINI_PROGRAM"
	LandingPageType_LANDING_PAGE_TYPE_STURNUS_IMAX                    LandingPageType = "LANDING_PAGE_TYPE_STURNUS_IMAX"
	LandingPageType_LANDING_PAGE_TYPE_TENCENT_NEWS_SOFT_ARTICLE       LandingPageType = "LANDING_PAGE_TYPE_TENCENT_NEWS_SOFT_ARTICLE"
	LandingPageType_LANDING_PAGE_TYPE_TENCENT_NEWS_LIVE               LandingPageType = "LANDING_PAGE_TYPE_TENCENT_NEWS_LIVE"
	LandingPageType_LANDING_PAGE_TYPE_XJ_OTT                          LandingPageType = "LANDING_PAGE_TYPE_XJ_OTT"
	LandingPageType_LANDING_PAGE_TYPE_YUEBAO_QUICKAPP                 LandingPageType = "LANDING_PAGE_TYPE_YUEBAO_QUICKAPP"
	LandingPageType_LANDING_PAGE_TYPE_YUEBAO_OFFICIAL_ACCOUNT_ARTICLE LandingPageType = "LANDING_PAGE_TYPE_YUEBAO_OFFICIAL_ACCOUNT_ARTICLE"
	LandingPageType_LANDING_PAGE_TYPE_XJ_QUICK                        LandingPageType = "LANDING_PAGE_TYPE_XJ_QUICK"
	LandingPageType_LANDING_PAGE_TYPE_QQ_MINI_GAME                    LandingPageType = "LANDING_PAGE_TYPE_QQ_MINI_GAME"
	LandingPageType_DEEP_LINK                                         LandingPageType = "DEEP_LINK"
	LandingPageType_H5                                                LandingPageType = "H5"
	LandingPageType_MINI_PROGRAM                                      LandingPageType = "MINI_PROGRAM"
	LandingPageType_UNIVERSAL_LINK                                    LandingPageType = "UNIVERSAL_LINK"
	LandingPageType_DEFAULT                                           LandingPageType = "DEFAULT"
	LandingPageType_DEEP_LINK_IOS                                     LandingPageType = "DEEP_LINK_IOS"
	LandingPageType_DEEP_LINK_ANDROID                                 LandingPageType = "DEEP_LINK_ANDROID"
)
