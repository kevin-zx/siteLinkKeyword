// some default filter
package filter

import (
	"github.com/kevin-zx/siteLinkKeyword/han"
	"github.com/kevin-zx/wordproperty"
	"strings"
)

var commStopWords = []string{"img", "上一篇", "下一篇", "更多", "首页", "全部", "上一页", "下一页", "登录", "设为首页", "下载", "关于我们", "加入收藏", "不限", "返回顶部", "立即下载", "最新", "联系我们", "查看详情", "详细", "查看更多", "退出", "新浪微博", "阅读全文", "QQ空间", "分享到", "详情", "关闭", "移动客户端", "注册", "取消", "推荐", "网站地图", "友情链接", "收藏", "意见反馈", "导航", "收藏本站", "腾讯微博", "回到顶部", "默认排序", "QQ好友", "问题库", "保存", "换肤", "评论", "新闻", "分享", "按更新时间", "搜索", "地形", "阅读", "换一换", "查看全文", "手机版", "按访问量", "复制链接", "订阅", "微博", "显示全部楼层", "下一章", "人人网", "忘记密码", "查看详细", "上一章", "视频", "详细介绍", "资讯", "全站导航", "猜你喜欢", "查看", "公告", "QQ", "专题", "热门", "第1页", "其他", "确定", "默认", "原创", "网站导航", "常见问题", "广告服务", "打印", "尾页", "收藏本页", "攻略", "最新发布", "地图", "百度贴吧", "百度", "论坛", "MORE", "加入我们", "综合", "专题首页", "最热门", "退出全屏", "网站首页", "引用", "English", "返回", "了解更多", "加关注", "UC浏览器", "中文", "帮助中心", "行业", "礼包", "next", "版权声明", "分享链接", "全屏显示", "prev", "刷新页面", "新闻中心", "更新时间", "其它", "品牌活动", "舆情中心", "国际World", "顶部", "收藏本文", "汽车Auto", "生命时报", "智能Smart", "艺术Art", "城市City", "历史记录", "招聘信息", "最新推荐", "纠错", "体育", "官网", "正文", "5秒注册", "产品", "转载", "我也说两句", "活跃榜", "刷新", "加入", "查看全部", "末页", "商务合作", "最近更新", "点击查看", "豆瓣网", "领取", "图片", "信誉", "返回列表", "免费下载", "进入频道", "more", "免责声明", "上页", "排行榜", "专区首页", "点击下载", "图片新闻", "动态", "了解详情", "专区", "生活", "查看全部评论", "在线客服", "全部评论", "全国", "主页", "客户服务", "文章", "进入专区", "关于本站", "关注", "Next", "电脑版", "音乐", "返回首页", "官方微博", "软件下载", "联系方式", "高速电信下载", "繁体版", "汽车", "网银登录", "本地高速下载", "简体中文", "查看全部留言", "综合排序", "最热", "举报", "收起", "报错", "人才招聘", "活动", "TOP", "补丁", "月排行", "下载帮助", "帮助", "专题推荐", "帐号登录", "回顶部", "去注册", "新闻资讯", "前进", "通知公告", "清空", "Menu", "More", "总排行", "客服中心", "免费注册", "关注更多", "全部分类", "我要提问", "在线咨询", "问答首页", "进入会员中心", "评测", "用户登录", "微博推荐", "收藏页面", "繁體中文", "售后服务", "条新消息", "条新评论", "热门频道", "房产问答", "菜单", "全文", "一键分享", "无障碍浏览", "娱乐", "服务", "原文链接", "分类", "个人中心", "回复", "发表评论", "软件发布", "按下载量", "热门推荐", "充值", "公司简介", "点评", "关于", "提交", "政策法规", "抢沙发", "页面", "历史", "反馈", "立即注册", "文档", "法律声明", "复制网址", "系统工具", "繁体", "电脑版下载", "社区", "软件", "益智", "概览", "海南", "新闻推荐", "频道", "品牌", "精品推荐", "点击进入", "浏览", "列表", "百科分类", "EN", "解决方案", "看过", "全部标签", "立即购买", "设为主页", "下页", "双人", "关注我们", "安卓版下载", "繁體", "我的账户", "QQ登录", "第一页", "下载客户端", "安卓软件", "更多分类", "问答", "进入官网", "旅游", "展开", "系列", "后退", "清空搜索记录", "广告合作", "综合统计", "借款分期", "立即申请", "申请与开通", "积分优惠", "基本资料", "区域论坛", "有更多新的消息", "电脑", "Android", "手机", "浙江", "二维码", "用户注册", "微信", "用户协议", "播放记录", "Prev", "注销", "特色市场", "购物车", "按人气排序", "隐私政策", "赛车竞速", "免费课程", "职业交流", "关闭窗口", "概述", "进入", "按时间", "直播电商", "退出登录", "全部问题", "排行", "中国", "提交建议", "吉林", "苹果下载", "互动", "投诉举报", "换一批", "商品", "综述", "信息公开", "投稿", "用户", "企业文化", "下载声明", "登录注册", "生活服务", "下载中心", "上一话", "诚聘英才", "新闻动态", "简介", "编辑精选", "工作动态", "服务条款", "底部", "主题", "客户端", "支付方式", "Close", "合作伙伴", "工具", "高速下载", "联系客服", "APP下载", "发展历程", "最新更新", "开始对比", "在线留言", "基本信息", "待解决", "今天", "全部主题", "设主页", "相关网站", "对比研究", "点击查看更多", "下载App", "关闭页面", "影音播放", "企业服务", "产品服务", "按时间排序", "简体", "业务办理", "0评论", "繁体中文", "官网首页", "安卓下载", "下载APP", "电信高速下载", "会员中心", "资料下载", "产品中心", "Search", "广告", "开始阅读", "删除", "支持", "女生", "去我的收藏夹", "投资者关系", "请登录", "技术支持", "查询", "优惠", "评论人参与", "安卓", "点击收藏kk", "作者", "片花预告", "看正片", "关闭菜单", "高级搜索", "字号增大", "字号减小", "软件教程", "问题反馈", "继续阅读", "专栏", "进入直播台", "分站", "我要分享", "换一张", "手机银行", "全部类型", "添加收藏", "网站介绍", "重新浏览", "Top", "投资理财", "Login", "立即登录", "聯系我們", "0626期", "手机查看", "合作厂商", "博客", "产品介绍", "资源下载", "要闻", "热点", "注册新用户", "确认", "执行禁封", "这里", "商城", "放到桌面", "热销产品", "将本站设为首页", "复制地址", "使用帮助", "精选品牌", "产品库", "按发布时间", "资料库", "返回目录", "网上银行登录", "用户中心", "广告联系", "阅读更多", "专题专栏", "最新资讯", "时尚", "百度新首页", "打印页面", "研究报告", "网摘", "立即咨询", "章节目录", "服务保障", "数据", "下载地址", "立即查看", "购物", "法律咨询", "CN", "Home", "亲子", "新闻评论", "个人", "策划", "马上登录", "全部清空", "TA的成就", "家长监护", "最新动态", "文化", "观看记录", "所有", "我知道了", "帐号", "经营", "邮箱", "软件专题", "文化艺术", "聊友分享", "分集剧情", "公众号", "年份", "0名", "更多内容", "清除", "政策解读", "iPhone", "热门标签", "最新文章", "注册账号", "相关阅读", "网友版", "国际", "资讯中心", "1个回答", "西安", "营业网点", "复制本文地址", "加载中", "附件", "和讯微博", "热门文章", "互动交流", "游戏图片", "标题", "知识产权", "进入专题", "close", "版权说明", "知道了", "ENGLISH", "售后咨询", "信用卡", "新闻列表", "当前位置", "入库时间", "立即验证", "首页Home", "立即兑换", "媒体报道", "公司动态", "知识", "设首页", "目录", "招聘", "网站声明", "返回登录", "查看原图", "在售", "Norway", "msg", "网站信息", "行业动态", "公司新闻", "好特播报", "更多分享", "购买", "浏览器", "咨询我", "去看看", "帐号设置", "历史通知", "去成长中心", "信息", "找回密码", "新用户注册", "不喜欢", "价格", "会员服务", "返回上一级", "手机软件", "分类页顶部广告", "查阅全文", "收藏到QQ书签", "点个赞", "评论页", "论坛转帖", "隐私保护", "收藏文章", "腾讯QQ", "加载更多", "看不清换一张", "分类页底部广告", "软件分类", "新游周刊", "咨询", "资源共享", "下载体验", "城市", "精彩图片", "加为收藏", "潮流推薦", "资源", "TA的主页", "一周", "分类目录", "上传", "登陆", "打印本页", "IOS", "俱乐部", "版权申明", "社交聊天", "精选", "全部软件", "国内", "联通高速下载", "站点地图", "首充号", "详细内容", "全屏", "更多资讯", "金融", "语言", "全部商品", "中文版", "过关", "EN中文简体", "PC", "加入收藏夹", "订阅说明", "链接", "专辑", "社会民生", "机构", "快速下载", "高智慧生命", "我要投稿", "点击阅读", "标签动态", "服务咨询", "下一图集", "报价", "收藏本网站", "欢迎投稿", "导购", "服务项目", "续充"}

// CommonStopWordFilter: the function used to filter website comm words like next, prev, etc.
func CommStopWordFilter(kim map[string]int) map[string]int {
	if len(kim) == 0 {
		return kim
	}
	for _, word := range commStopWords {
		delete(kim, word)
	}
	return kim
}

var nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "-", ")", "(", "（", "）"}

// Some time in some coarse site, they'll use productA1, productA2, productA3 to show productA's
// different pictures or contents.
// So TailNumFilter function is for this situation to clear insignificant tail.
func TailNumFilter(keyM map[string]int) map[string]int {
	tmpKm := make(map[string]int)
	for k, count := range keyM {
		replace := false
		key := k
		for true {
			containsNum := false
			for _, num := range nums {
				if strings.HasSuffix(key, num) {
					key = strings.ReplaceAll(key, num, "")
					containsNum = true
					replace = true
				}
			}
			// if not has any suffix nums or some specified symbols, will break the loop
			if !containsNum {
				break
			}
		}
		if replace {
			delete(keyM, k)
			tmpKm[key] = count
		}
	}
	// I just know in a map for loop, delete operation is safety, but don't know add operations is whether safety.
	// So I do this operation out of the loop in case some exceptions.
	for k, c := range tmpKm {
		keyM[k] = c
	}
	return keyM
}

// This filter will judge a word, whether a regular Chinese word, if not, then will delete it.
func FormalWordFilter(keyM map[string]int) map[string]int {
	for k, _ := range keyM {
		if !han.IsHan(k) {
			delete(keyM, k)
		}
	}
	return keyM
}

var restOneWords = []string{"微信", "qq", "电话", "QQ"}

// some word is not comm but also isn't industry word, like company name + qq, company name + phone, etc.
// we take the tactic that reset this word count to one in this function.
func RestOneFilter(keyM map[string]int) map[string]int {
	for k, _ := range keyM {
		for _, word := range restOneWords {
			if strings.Contains(k, word) {
				keyM[k] = 1
				break
			}
		}
	}
	return keyM
}

var punctuations = []string{"「", "」", "•", "　", "«", "›", ">", "»", "|", "(", ")", "（", "）", "\t", "\r", "\n", " ", "©", "·", "=", "-", "，", "。", "、", "；", "’", "【", "】", "、", "`", "！", "@", "#", "￥", "%", "…", "…", "&", "×", "—", "—", "《", "》", "？", "：", "”", "“", "{", "}", "‘", "|", "～", "+", ",", "/", ";", "'", "[", "]", "\\", "`", "!", "@", "#", "$", "%", "^", "&", "*", "_", "+", "<", ">", "?", ":", "\"", "{", "}", "~"}

// ClearPunctuationsFilter: In a typical sense, a word does not have any punctuation (except c#, .net, etc.).
// So this function used to clear those punctuations.
func ClearPunctuationsFilter(keyM map[string]int) map[string]int {
	tmpKm := make(map[string]int)
	for k, c := range keyM {
		r := false
		tk := k
		for _, punctuation := range punctuations {
			if strings.Contains(tk, punctuation) {
				tk = strings.ReplaceAll(tk, punctuation, "")
				r = true
			}
		}

		if r {
			tmpKm[tk] = c
			delete(keyM, k)
		}
	}

	for k, c := range tmpKm {
		keyM[k] = c
	}
	return keyM
}

var trimWords = []string{" ", " ", "\uE604", "\uE613", "\r", "\n", "\t"}

// this function is straightforward, trim words, but on the website, we will find some magic blank's symbols,
// there collect some symbols for trim
func TrimFilter(keyM map[string]int) map[string]int {
	tmpKm := make(map[string]int)
	for k, c := range keyM {
		kt := k
		m := false
		for _, trimWord := range trimWords {
			if strings.Contains(k, trimWord) {
				kt = strings.ReplaceAll(kt, trimWord, "")
				m = true
			}
		}
		if m {
			tmpKm[kt] = c
			delete(keyM, k)
		}
	}
	for k, c := range tmpKm {
		keyM[k] = c
	}
	return keyM
}

// if a word len greater than 10, maybe the word isn't a word, it's a sentence
func KeywordLenFilter(keyM map[string]int) map[string]int {
	for k := range keyM {
		if len(strings.Split(k, "")) >= 10 {
			keyM[k] = 1
		}
	}
	return keyM
}

//  filter some environment word
func EnvFilter(keyM map[string]int) map[string]int {
	for k := range keyM {
		if ok, _ := wordproperty.EnvWordProperty(k); ok {
			delete(keyM, k)
		}
	}
	return keyM
}
