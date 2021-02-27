package tiktokdl

import (
	"os"

	"github.com/siongui/tiktokgo"
)

// DownloadAvatar downloads user's larger avatar (profile) photo.
func DownloadAvatar(user tiktokgo.TiktokUser) (err error) {
	err = CreateDirIfNotExist(UserDir(user))
	if err != nil {
		return
	}

	avatarpath := UserAvatarFilePath(user)
	if _, err := os.Stat(avatarpath); os.IsNotExist(err) {
		return SaveItemBytes(user.AvatarLarger, avatarpath, nil, nil)
	} else {
		if err != nil {
			return err
		}
	}
	return
}

// DownloadItem downloads user's tiktok video.
func DownloadItem(item tiktokgo.TiktokItem) (err error) {
	cookies, err := tiktokgo.GetCookies()
	if err != nil {
		return
	}

	err = CreateDirIfNotExist(UserDir(item.Author))
	if err != nil {
		return
	}

	itempath := UserItemFilePath(item)
	if _, err := os.Stat(itempath); os.IsNotExist(err) {
		return SaveItemBytes(item.Video.PlayAddr, itempath, tiktokgo.GetHeaders(), cookies)
		//println(item.Video.PlayAddr)
		//println(item.Video.DownloadAddr)
		//println(itempath)
	} else {
		if err != nil {
			return err
		}
	}
	return
}

// DownloadUserPageNextData downloads larger user profile pic and video items in
// the __NEXT_DATA__ of user page HTML.
func DownloadUserPageNextData(username string) (err error) {
	nd, err := tiktokgo.GetWebUserPageNextData(username)
	if err != nil {
		return
	}

	userinfo := nd.Props.PageProps.UserInfo
	err = DownloadAvatar(userinfo.User)
	if err != nil {
		return
	}

	items := nd.Props.PageProps.Items
	for _, item := range items {
		err = DownloadItem(item)
		if err != nil {
			return
		}
	}

	return
}
