/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-02-05 20:12
 * description :
 * history :
 */
package member

import (
	"go2o/core/infrastructure/domain"
)

var (
	ErrInvitationCode *domain.DomainError = domain.NewDomainError(
		"member_err_invation_code", "CODE:1011,邀请码错误")

	ErrRegOff *domain.DomainError = domain.NewDomainError(
		"err_reg_off", "CODE:1011,系统未开放注册")

	ErrRegMustInvitation *domain.DomainError = domain.NewDomainError(
		"err_reg_must_invitation", "CODE:1011,系统只允许邀请注册")

	ErrRegOffInvitation *domain.DomainError = domain.NewDomainError(
		"err_reg_off_invitation", "CODE:1011,系统关闭邀请注册")

	ErrSessionTimeout *domain.DomainError = domain.NewDomainError(
		"member_session_time_out", "会员会话超时")

	ErrCheckCodeError *domain.DomainError = domain.NewDomainError(
		"err_member_check_code_err", "校验码不正确")

	ErrCheckCodeExpires *domain.DomainError = domain.NewDomainError(
		"err_member_check_code_expires", "校验码已失效")

	ErrInvalidSession *domain.DomainError = domain.NewDomainError(
		"member_invalid_session", "异常会话")

	ErrPwdLength *domain.DomainError = domain.NewDomainError(
		"err_member_pwd_length", "密码至少包含6个字符")

	ErrNoSuchDeliverAddress *domain.DomainError = domain.NewDomainError(
		"member_no_such_deliver_address", "配送地址错误")

	ErrLevelUsed *domain.DomainError = domain.NewDomainError(
		"member_level_used", "此等级已被会员使用")

	ErrLevelRequireExp *domain.DomainError = domain.NewDomainError(
		"member_level_require_exp", "所需经验值必须大于%d")

	ErrNoSuchMember *domain.DomainError = domain.NewDomainError(
		"member_no_such_member", "会员不存在")

	ErrDeliverAddressLen *domain.DomainError = domain.NewDomainError(
		"err_deliver_address_len", "请填写详细的配送地址")

	ErrDeliverContactPersonName *domain.DomainError = domain.NewDomainError(
		"err_deliver_contact_person_name", "收货人不正确")

	ErrDeliverContactPhone *domain.DomainError = domain.NewDomainError(
		"err_deliver_phone_is_null", "联系人电话有误")

	ErrNotSetArea *domain.DomainError = domain.NewDomainError(
		"err_not_set_area", "地址不正确")

	ErrNoSuchBankInfo *domain.DomainError = domain.NewDomainError(
		"err_no_such_bank_info", "请完善银行卡信息")

	ErrBankInfoLocked *domain.DomainError = domain.NewDomainError(
		"err_bank_info_locked", "银行卡信息已锁定,无法更改")

	ErrBankInfoNoYetSet *domain.DomainError = domain.NewDomainError(
		"err_bank_info_no_yet_set", "银行卡信息尚未设置")

	ErrIncorrectAmount *domain.DomainError = domain.NewDomainError(
		"err_balance_amount", "金额错误")

	ErrOutOfBalance *domain.DomainError = domain.NewDomainError(
		"err_out_of_balance", "超出金额")

	ErrUsrLength *domain.DomainError = domain.NewDomainError(
		"err_user_length", "用户名至少6位",
	)

	ErrUsrValidErr *domain.DomainError = domain.NewDomainError(
		"err_user_valid_err", "用户名为6位以上字符和数字的组合")

	ErrSameUsr *domain.DomainError = domain.NewDomainError(
		"err_same_usr", "用户名与原来相同")

	ErrUsrExist *domain.DomainError = domain.NewDomainError(
		"err_member_usr_exist", "用户名已存在")

	ErrNilNickName *domain.DomainError = domain.NewDomainError(
		"err_member_nil_nick_name", "昵称不能为空")

	ErrEmailValidErr *domain.DomainError = domain.NewDomainError(
		"err_member_email_valid_err", "邮箱不正确")

	ErrPhoneValidErr *domain.DomainError = domain.NewDomainError(
		"err_member_phone_valid_err", "手机号码不正确")

	ErrPhoneHasBind *domain.DomainError = domain.NewDomainError(
		"err_member_phone_has_bind", "手机号码已经绑定")

	ErrMissingPhone *domain.DomainError = domain.NewDomainError(
		"err_member_missing_phone", "请填写手机号码")

	ErrMissingIM *domain.DomainError = domain.NewDomainError(
		"err_member_missing_im", "请填写IM")

	ErrBadPhoneFormat *domain.DomainError = domain.NewDomainError(
		"err_bad_phone_format", "手机号码不正确")

	ErrQqValidErr *domain.DomainError = domain.NewDomainError(
		"err_qq_valid_err", "QQ号码不正确")

	ErrNotSetTradePwd *domain.DomainError = domain.NewDomainError(
		"err_not_set_tarde_pwd", "交易密码未设置")

	ErrIncorrectTradePwd *domain.DomainError = domain.NewDomainError(
		"err_incorrect_tarde_pwd", "交易密码错误")

	ErrNotEnoughAmount *domain.DomainError = domain.NewDomainError(
		"err_not_enough_amount", "金额不足")

	ErrNotSupportTransfer *domain.DomainError = domain.NewDomainError(
		"err_not_support_transfer", "不支持的转账方式")

	ErrMissingTrustedInfo *domain.DomainError = domain.NewDomainError(
		"err_missing_trusted_info", "实名认证信息不完整")

	ErrFavored *domain.DomainError = domain.NewDomainError(
		"err_favored", "已经收藏过了")

	ErrAccountBalanceNotEnough *domain.DomainError = domain.NewDomainError(
		"err_account_balance_not_enough ", "账户余额不足")
)
