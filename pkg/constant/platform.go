package constant

const (
	//Platform ID
	IOSPlatformID        = 1
	AndroidPlatformID    = 2
	WindowsPlatformID    = 3
	OSXPlatformID        = 4
	WebPlatformID        = 5
	MiniWebPlatformID    = 6
	LinuxPlatformID      = 7
	AndroidPadPlatformID = 8
	IPadPlatformID       = 9

	//Platform string match to Platform ID
	IOSPlatformStr        = "IOS"
	AndroidPlatformStr    = "Android"
	WindowsPlatformStr    = "Windows"
	OSXPlatformStr        = "OSX"
	WebPlatformStr        = "Web"
	MiniWebPlatformStr    = "MiniWeb"
	LinuxPlatformStr      = "Linux"
	AndroidPadPlatformStr = "APad"
	IPadPlatformStr       = "IPad"
)

var PlatformID2Name = map[int]string{
	IOSPlatformID:        IOSPlatformStr,
	AndroidPlatformID:    AndroidPlatformStr,
	WindowsPlatformID:    WindowsPlatformStr,
	OSXPlatformID:        OSXPlatformStr,
	WebPlatformID:        WebPlatformStr,
	MiniWebPlatformID:    MiniWebPlatformStr,
	LinuxPlatformID:      LinuxPlatformStr,
	AndroidPadPlatformID: AndroidPadPlatformStr,
	IPadPlatformID:       IPadPlatformStr,
}

func PlatformIDToName(n int) string {
	return PlatformID2Name[n]
}
