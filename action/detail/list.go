package detail

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func List(ctx *gin.Context) {
	data := `{
	"code": 0,
	"body": {
			"bookInfo": {
					"currentTime": 1606273076,
					"bookId": 11679916,
					"bookName": "你还未嫁我怎敢老（漫画版）",
					"magazineId": 0,
					"feeUnit": 20,
					"lastChapterTime": "2019-12-27 16:45:44",
					"publisher": "",
					"completeState": "Y",
					"circleId": "book_11679916",
					"picUrl": "http://book.img.ireader.com/idc_1/m_1,w_300,h_400,f_webp/ad30df8e/group6/M00/63/DE/CmRaIVtFsmaEEMweAAAAAEGgDHM015296232.jpg?v=9F0QVK9L&t=CmQUN16IvGU.",
					"headPic": "http://book.img.ireader.com/idc_1/f_webp/1595746c/group6/M00/6F/72/CmRaIVt_HxKEfc9eAAAAABou9bA338562448.jpg?v=A_kGYtNY&t=CmRaIVvyKP0.",
					"wordCount": "590字",
					"isbn": "",
					"bookType": 2,
					"icon": "漫画",
					"limitInfo": [],
					"voteNum": 505,
					"priceInfo": {
							"delPrice": "49阅饼/话",
							"activePrice": "免费",
							"isFree": true,
							"discount": 0,
							"timeType": 0,
							"tag": "",
							"label": "",
							"promotionMark": 64
					},
					"author": "掌阅独家",
					"desc": "猝不及防的初遇，她懵懂中一见钟情。岁月静好时相处，他无意间怦然心动。六年分离，彼此渐行渐远，仍然小心翼翼守护心中挚爱。阔别重逢，苏沫已不是富家千金，收敛锋芒努力生活之中依旧对他心心念念；陆景炎也不再少年无忧，冷漠绝情多番拒绝之下难抑为她寸心如狂。花还未落，树怎敢死；你还未嫁，我怎敢老。",
					"categorys": [
							{
									"id": 417,
									"name": "少女漫画"
							},
							{
									"id": 1366,
									"name": "少女漫画"
							},
							{
									"id": 432,
									"name": "恋爱物语"
							},
							{
									"id": 422,
									"name": "小说改编"
							}
					],
					"fromSource": "掌阅漫画",
					"authorList": [
							{
									"id": 794990,
									"circleId": "author_794990",
									"name": "掌阅独家",
									"type": "",
									"usr": "i1732831912",
									"url": "http://ah2.zhangyue.com/zyuc/homepage/home/index?p1=VyVHBVzzuEEDACztLBJcEYD6&p2=119042&p3=17150003&p4=501603&p5=16&p6=IJIGABBIIACBCHHFJFJE&p7=__624150017921616&p9=46009&p12=&p16=vivo+V3M+A&p21=31303&p22=5.1&p25=7150001&p26=22&usr=i2995737697&rgt=7&zyeid=b32ddea143274bbaa590b89b18a85e2b&visitorName=i1732831912"
							},
							{
									"id": 818036,
									"circleId": "",
									"name": "LeeLee啊木木绘制",
									"type": "",
									"usr": "i1732901461",
									"url": "http://ah2.zhangyue.com/zyuc/homepage/home/index?p1=VyVHBVzzuEEDACztLBJcEYD6&p2=119042&p3=17150003&p4=501603&p5=16&p6=IJIGABBIIACBCHHFJFJE&p7=__624150017921616&p9=46009&p12=&p16=vivo+V3M+A&p21=31303&p22=5.1&p25=7150001&p26=22&usr=i2995737697&rgt=7&zyeid=b32ddea143274bbaa590b89b18a85e2b&visitorName=i1732901461"
							}
					],
					"tagInfo": [
							{
									"name": "霸道总裁",
									"url": "http://ah2.zhangyue.com/zybk/tag/searchbook?tag=%E9%9C%B8%E9%81%93%E6%80%BB%E8%A3%81&sort=6"
							},
							{
									"name": "现代都市",
									"url": "http://ah2.zhangyue.com/zybk/tag/searchbook?tag=%E7%8E%B0%E4%BB%A3%E9%83%BD%E5%B8%82&sort=6"
							},
							{
									"name": "少女动漫",
									"url": "http://ah2.zhangyue.com/zybk/tag/searchbook?tag=%E5%B0%91%E5%A5%B3%E5%8A%A8%E6%BC%AB&sort=6"
							},
							{
									"name": "文改漫",
									"url": "http://ah2.zhangyue.com/zybk/tag/searchbook?tag=%E6%96%87%E6%94%B9%E6%BC%AB&sort=6"
							}
					],
					"driveTag": [],
					"putAwayDate": "2018-07-11",
					"listenBookInfo": [],
					"cartoonBookInfo": [],
					"breakPayInfo": "",
					"recommend": {},
					"responsibleEditor": "",
					"attribute": {
							"weekDownPv": "1398",
							"weekDownPvDesc": "在读",
							"likeNum": "2.2万",
							"likeNumDesc": "点赞",
							"star": 8.8,
							"starStyle": 9,
							"popularity": "1412.4万",
							"popularityDesc": "人气",
							"popularityNum": 14123829,
							"fansNum": "7756",
							"fansNumDesc": "粉丝"
					},
					"cpId": "2446",
					"cpName": "掌阅漫画",
					"size": null,
					"showEpubSerialWarning": false,
					"lastSupportVersion": 700003
			},
			"chaperInfo": {
					"chapterName": "31 - 最后",
					"chapterNum": 31,
					"orgStatus": "已完结"
			},
			"iconInfo": {
					"markInfo": {
							"isMark": true,
							"data": {
									"id": 0,
									"vote": 0,
									"status": false
							}
					},
					"bookListInfo": {
							"isBookList": true,
							"data": {
									"id": 11679916,
									"name": "你还未嫁我怎敢老（漫画版）",
									"author": "掌阅独家",
									"picUrl": "http://book.img.ireader.com/idc_1/m_1,w_300,h_400,f_webp/ad30df8e/group6/M00/63/DE/CmRaIVtFsmaEEMweAAAAAEGgDHM015296232.jpg?v=9F0QVK9L&t=CmQUN16IvGU.",
									"status": false
							}
					},
					"cartInfo": {
							"isCart": false,
							"data": {
									"status": 0,
									"isNewCart": 0
							},
							"url": ""
					},
					"rewardInfo": {
							"isReward": true,
							"data": {
									"isActivity": false
							},
							"url": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Ticket.Index&circleId=book_11679916&bookId=11679916"
					},
					"shareInfo": {
							"isShare": true,
							"data": {
									"Action": "share",
									"Data": {
											"summary": "世界很美,旅途很长……带上一本好书《你还未嫁我怎敢老（漫画版）》,和我一起聆听时光",
											"author": "掌阅独家",
											"title": "发现一本好书《你还未嫁我怎敢老（漫画版）》",
											"picUrl": "http://book.img.ireader.com/idc_1/m_1,w_300,h_400/739ed8ed/group6/M00/63/DE/CmRaIVtFsmaEEMweAAAAAEGgDHM015296232.jpg?v=9F0QVK9L&t=CmQUN16IvGU.",
											"attract": {
													"type": "book",
													"attr": {
															"id": "11679916"
													},
													"pos": "bookdetail"
											},
											"Speaker": null,
											"shareType": "none",
											"url": "http://ah2.zhangyue.com/zybook3/u/p/api.php?Act=weixin&bid=11679916&shareusr=i2995737697&p2=119042&p3=17150003&fid=41"
									}
							}
					}
			},
			"vipZoneInfo": {
					"desc": "",
					"sVipDesc": "",
					"statusDesc": "",
					"url": "",
					"isSVipUser": false
			},
			"svipZoneInfo": {},
			"commentList": {
					"list": [
							{
									"id": 176171242,
									"liked": false,
									"name": "",
									"likeNum": 6,
									"replyNum": 0,
									"isAuthor": false,
									"isAdmin": false,
									"vote": 5,
									"usr": "i792705832",
									"nick": "喜虐喜冷",
									"isVip": false,
									"userVipStatus": "1",
									"level": 11,
									"isV": false,
									"vType": null,
									"source": "comment",
									"ts": "7月16日",
									"content": "很好看画的也很好，我知道男主是爱着女主的只是因为那些原因暂时不可以说明彼此的心意，女主小时候好可爱啊，男主高高冷冷的但是对女主可以看出他的温柔。",
									"avatar": "http://ibk.ugc.corp3g.cn/idc_1/m_4,w_60,h_60,q_100,f_webp/c0b45da6/group6/M00/AE/92/CmQUNlgdo76EJtWpAAAAAFD5IVc970342759.jpg?v=qb4ZPyip&t=CmQUNlyZgcg.",
									"avatarFrame": "http://book.img.ireader.com/idc_1/m_1,w_300,h_400/9fce3127/group6/M00/D3/9D/CmQUNlpyyFGEXC12AAAAAItBnHU940272724.png?v=AeJNNiNy&t=CmQUNlpyyFE.",
									"spaceUrl": "http://ah2.zhangyue.com/zybook3/app/app.php?ca=User_Space.Index&un=i792705832&mod=bookDetail",
									"url": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Topic.Detail&key=BCM176171242&topicId=176171242&circleId=book_11679916&pk=client_bkCmt",
									"sUrl": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Group.Show&key=BGROUP_book_11679916&id=book_11679916&topicType=comment&cache=no&pk=client_bkCmt_more"
							},
							{
									"id": 172715753,
									"liked": false,
									"name": "",
									"likeNum": 28,
									"replyNum": 0,
									"isAuthor": false,
									"isAdmin": false,
									"vote": 5,
									"usr": "i2409074829",
									"nick": "i24****4829",
									"isVip": false,
									"userVipStatus": "1",
									"level": 4,
									"isV": false,
									"vType": null,
									"source": "comment",
									"ts": "5月11日",
									"content": "喜欢，它不是口中的天花乱坠，而是附有感情的真实情感，同样，我对这部作品也是一样的，微笑婉转",
									"avatar": "https://fs-uc-nearme-com-cn.oss-cn-hangzhou.aliyuncs.com/default.png",
									"avatarFrame": "http://book.img.ireader.com/idc_1/m_1,w_300,h_400/9fce3127/group6/M00/D3/9D/CmQUNlpyyFGEXC12AAAAAItBnHU940272724.png?v=AeJNNiNy&t=CmQUNlpyyFE.",
									"spaceUrl": "http://ah2.zhangyue.com/zybook3/app/app.php?ca=User_Space.Index&un=i2409074829&mod=bookDetail",
									"url": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Topic.Detail&key=BCM172715753&topicId=172715753&circleId=book_11679916&pk=client_bkCmt",
									"sUrl": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Group.Show&key=BGROUP_book_11679916&id=book_11679916&topicType=comment&cache=no&pk=client_bkCmt_more"
							},
							{
									"id": 144404615,
									"liked": false,
									"name": "",
									"likeNum": 15,
									"replyNum": 0,
									"isAuthor": false,
									"isAdmin": false,
									"vote": 5,
									"usr": "i1847817210",
									"nick": "梦天使楚兮霖",
									"isVip": false,
									"userVipStatus": "1",
									"level": 5,
									"isV": false,
									"vType": null,
									"source": "comment",
									"ts": "2019年3月9日",
									"content": "求更新版本！求问心无愧就好😄！～支持作者大大更新的人点赞👍或者是给个群号让我们知道你的更新版本不同啊😱(๑•̀ㅂ•́)ง<zyemot>难过_2_44_48_48</zyemot>",
									"avatar": "http://ibk.ugc.corp3g.cn/idc_1/m_4,w_60,h_60,q_100,f_webp/345d8f3f/group61/M00/47/93/CmQUOVw8MEmEYXPXAAAAAFNIRq0937362732.jpg?v=aLv6qkUA&t=CmQUOV2xyyc.",
									"avatarFrame": "http://book.img.ireader.com/idc_1/m_1,w_300,h_400/9fce3127/group6/M00/D3/9D/CmQUNlpyyFGEXC12AAAAAItBnHU940272724.png?v=AeJNNiNy&t=CmQUNlpyyFE.",
									"spaceUrl": "http://ah2.zhangyue.com/zybook3/app/app.php?ca=User_Space.Index&un=i1847817210&mod=bookDetail",
									"url": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Topic.Detail&key=BCM144404615&topicId=144404615&circleId=book_11679916&pk=client_bkCmt",
									"sUrl": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Group.Show&key=BGROUP_book_11679916&id=book_11679916&topicType=comment&cache=no&pk=client_bkCmt_more"
							},
							{
									"id": 141763978,
									"liked": false,
									"name": "",
									"likeNum": 101,
									"replyNum": 0,
									"isAuthor": false,
									"isAdmin": false,
									"vote": 5,
									"usr": "i1019127157",
									"nick": "初心不变",
									"isVip": false,
									"userVipStatus": "1",
									"level": 4,
									"isV": false,
									"vType": null,
									"source": "comment",
									"ts": "2019年2月2日",
									"content": "突然觉得自己和女主很像，当年也是追着一个男生到处跑。但他和男主却一点都不像",
									"avatar": "http://ibk.ugc.corp3g.cn/idc_1/m_4,w_60,h_60,q_100,f_webp/43c436a0/group6/M00/EE/C8/CmQUNlss3DWEWnQuAAAAAKtwq3E434793098.jpg?v=_924UTCm&t=CmQUNlss3DU.",
									"avatarFrame": "http://book.img.ireader.com/idc_1/m_1,w_300,h_400/9fce3127/group6/M00/D3/9D/CmQUNlpyyFGEXC12AAAAAItBnHU940272724.png?v=AeJNNiNy&t=CmQUNlpyyFE.",
									"spaceUrl": "http://ah2.zhangyue.com/zybook3/app/app.php?ca=User_Space.Index&un=i1019127157&mod=bookDetail",
									"url": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Topic.Detail&key=BCM141763978&topicId=141763978&circleId=book_11679916&pk=client_bkCmt",
									"sUrl": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Group.Show&key=BGROUP_book_11679916&id=book_11679916&topicType=comment&cache=no&pk=client_bkCmt_more"
							},
							{
									"id": 178551895,
									"liked": false,
									"name": "",
									"likeNum": 4,
									"replyNum": 0,
									"isAuthor": false,
									"isAdmin": false,
									"vote": 4,
									"usr": "i312035954",
									"nick": "童糕i",
									"isVip": false,
									"userVipStatus": "1",
									"level": 5,
									"isV": false,
									"vType": null,
									"source": "comment",
									"ts": "8月27日",
									"content": "刚刚看完了这个，一共也就30多话，很容易就看完，但是他后面就不更新了，有些遗憾。好像还有小说版的，我去看看小说版的吧，但不过真的希望他继续更新。剧情嘛，这个男主他就是，典型的钢铁直男，然后做法有些让人生气。女主呢性格一直很好，但是我看小时候是长发，为什么长大就是短发呢？短发不好看。男二也是又美又帅<zyemot>坏笑_2_57_48_48</zyemot>那个什么若微，她表里不一，表面性格好，内心却<zyemot>恶心_2_49_48_48</zyemot>。总体来说还是挺喜欢的(๑•̀ㅂ•́)ง",
									"avatar": "http://ibk.ugc.corp3g.cn/idc_1/m_4,w_60,h_60,q_100,f_webp/1a6de003/group6/M00/0B/83/CmRablrK9U6EL9CyAAAAABnsUK8995240251.jpg?v=g1ogb7Br&t=CmQUNl9KwGc.",
									"avatarFrame": "http://book.img.ireader.com/idc_1/m_1,w_300,h_400/9fce3127/group6/M00/D3/9D/CmQUNlpyyFGEXC12AAAAAItBnHU940272724.png?v=AeJNNiNy&t=CmQUNlpyyFE.",
									"spaceUrl": "http://ah2.zhangyue.com/zybook3/app/app.php?ca=User_Space.Index&un=i312035954&mod=bookDetail",
									"url": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Topic.Detail&key=BCM178551895&topicId=178551895&circleId=book_11679916&pk=client_bkCmt",
									"sUrl": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Group.Show&key=BGROUP_book_11679916&id=book_11679916&topicType=comment&cache=no&pk=client_bkCmt_more"
							}
					],
					"circleInfo": {
							"topicNum": 371,
							"replyNum": 68,
							"totalNum": "439",
							"fansNum": "7756",
							"circleId": "book_11679916",
							"url": "http://ah2.zhangyue.com/zysns/app/app.php?ca=Group.Show&id=book_11679916&topicType=all&from=bkdetail&cache=no"
					}
			},
			"commentSwitch": 1,
			"vipBuyBookInfo": {
					"isVipBook": 1,
					"isVipUser": 0,
					"isComicBook": 1,
					"vBuy": 0
			},
			"gift": [],
			"comicThumbs": {},
			"sharePreferential": {
					"btnDesc": "分享好书给小伙伴",
					"isShowBtn": true
			},
			"redPacket": {
					"show": 0,
					"url": "http://ah2.zhangyue.com/zybk/api/detail/redpacket?bid=11679916&resId=11679916&rpId="
			},
			"disclaimer": " 本书数字版权由“掌阅漫画”提供，并由其授权掌阅科技股份有限公司制作发行，若书中含有不良信息，请书友积极告知客服。",
			"wonderfulNote": {},
			"usrRecCard": {
					"usr": "i1732831912",
					"authorId": 794990,
					"circleId": "author_794990",
					"authorName": "i17****1912",
					"authorIcon": "http://ibk.ugc.corp3g.cn/idc_1/m_1,w_300,h_400/3f36f8e6/group61/M00/98/1D/CmQUOF4X-wmEW5BCAAAAAAT7nDA300529131.png?v=eb8ydy91&t=CmQUOF4X-wk.",
					"authorRankTag": "",
					"userTag": "author",
					"userVInfo": {
							"tag": false,
							"intro": ""
					},
					"userTagName": "作者",
					"identityIntro": "",
					"isShowFollow": true,
					"verification": {},
					"isFollow": 0,
					"isSupportLocal": 1,
					"followUrl": "http://ah2.zhangyue.com/zyuc/user/follow?userName=i1732831912&act=follow",
					"url": "http://ah2.zhangyue.com/zyuc/homepage/home/index?visitorName=i1732831912",
					"authorMedal": {
							"list": [
									"http://book.img.ireader.com/idc_1/group6/M00/24/89/CmRabltegeeEcfjXAAAAAE7d0Bg621925557.png?v=i5akg8r_&t=CmRabltegec.",
									"http://book.img.ireader.com/idc_1/group6/M00/24/88/CmRae1tegraETmSXAAAAAB8tsVg103879981.png?v=BsyuvNNJ&t=CmRae1tegrY.",
									"http://book.img.ireader.com/idc_1/group6/M00/24/89/CmRae1tehcqEOimmAAAAAMkhsQE025766584.png?v=5EruwzDF&t=CmRae1tehco.",
									"http://book.img.ireader.com/idc_1/group6/M00/F7/63/CmQUNltegVmECqMEAAAAAJel4oI449130029.png?v=dliKQTrb&t=CmQUNltegVk."
							],
							"text": "荣获4枚勋章",
							"url": "http://ah2.zhangyue.com/zyuc/profile/medal/index?visitorName=i1732831912",
							"canJump": 1
					},
					"fansCountText": "",
					"wordCountText": "累计2736话"
			},
			"buttonInfo": [
					{
							"btnText": "在线阅读",
							"encStr": "dBHOPzJBzVduu8e0%2BlffQqGQCjDAVSIxwr%2BRrKuSVwol%2FGSOCeU%2FJKQGacv9OvuDSIfFzZ9nJyQIfSevlcbJ5QEh2RWJndIO7lfEJ9cFZis%3D",
							"type": "free",
							"cmd": {
									"Action": "onlineReader",
									"Data": {
											"Charging": {
													"FeeType": 0,
													"OrderUrl": "http://ah2.zhangyue.com/zytc/public/index.php?ca=Order.Create&bid=11679916&cid=1&vBuy=0&projectSource=zybook3&pk=ATUHORALL",
													"Price": 0
											},
											"DownloadInfo": {
													"ChapterId": 1,
													"DownloadUrl": "http://ah2.zhangyue.com/r/enc_dl?downInfo=ZPLBHvmGGzyw4ssvbcwk5Hz8HpXbRKsbB4pzuMEMJ73s_1NFhXIc6mjTdEPP-sHpnIdZpZRR38dPZTBO4-F_NyTav13j793TIsS-YcRF-7W9jV1OEO1CaIkQf6BrlKoBjUNY7MiaWglF28Kw2jS1l7hMkZQYjcXuDKNEMzlloDk",
													"Ebk3DownloadUrl": "http://ah2.zhangyue.com/r/enc_dl?downInfo=ZPLBHvmGGzyw4ssvbcwk5Hz8HpXbRKsbB4pzuMEMJ73s_1NFhXIc6mjTdEPP-sHpnIdZpZRR38dPZTBO4-F_NyTav13j793TIsS-YcRF-7W9jV1OEO1CaIkQf6BrlKoBjUNY7MiaWglF28Kw2jS1l7hMkZQYjcXuDKNEMzlloDk",
													"FeeUnit": 20,
													"FileId": "11679916",
													"FileName": "《你还未嫁我怎敢老[漫画]》.epub",
													"FileSize": 10000000,
													"Type": "2",
													"Version": "2",
													"OrderId": "B0125323729944187904"
											},
											"bookCatalog": {
													"type": 2,
													"bookId": 11679916,
													"downloadType": 0,
													"data": {
															"orderId": 0,
															"genreId": 0,
															"genreDate": "",
															"genreName": ""
													},
													"isEpubSerialize": 0,
													"bookType": "2",
													"relResource": [
															{
																	"status": null,
																	"book_id": null,
																	"type": null
															}
													],
													"relBookId": "0"
											}
									}
							}
					}
			]
	},
	"msg": "success"
}`
	ctx.JSON(http.StatusOK, data)
}
