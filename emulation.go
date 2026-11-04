package cdp

import (
	"context"
	"strings"

	"github.com/gospider007/gson"
)

// 设置userAgent
func (obj *WebSock) EmulationSetUserAgentOverride(preCtx context.Context, userAgent string, acceptLanguage string, platform string, userAgentData UserAgentData) (RecvData, error) {
	// log.Print(userAgentData.FullVersionList)
	// log.Print(userAgentData.UaFullVersion)
	params := map[string]any{
		"userAgentMetadata": userAgentData,
		// "userAgentMetadata": UserAgentData{
		// 	Brands:          userAgentData.Brands,
		// 	FullVersionList: userAgentData.FullVersionList,
		// 	UaFullVersion:   userAgentData.UaFullVersion,
		// 	Platform:        userAgentData.Platform,
		// 	PlatformVersion: userAgentData.PlatformVersion,
		// 	Architecture:    userAgentData.Architecture,
		// 	Model:           userAgentData.Model,
		// 	Mobile:          userAgentData.Mobile,
		// 	Bitness:         userAgentData.Bitness,
		// 	Wow64:           userAgentData.Wow64,
		// 	FormFactors:     userAgentData.FormFactors,
		// },
	}

	if platform != "" {
		params["platform"] = platform
	}
	if userAgent != "" {
		params["userAgent"] = userAgent
	}
	if acceptLanguage != "" {
		params["acceptLanguage"] = acceptLanguage
	}
	return obj.send(preCtx, commend{
		Method: "Emulation.setUserAgentOverride",
		Params: params,
	})
}

type Viewport struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
type Device struct {
	UserAgent         string   `json:"user_agent"`
	Viewport          Viewport `json:"viewport"`            //浏览器宽度和高度
	DeviceScaleFactor float64  `json:"device_scale_factor"` //1 → 普通屏，2 → Retina，3 → 高密度手机屏
	IsMobile          bool     `json:"is_mobile"`
	HasTouch          bool     `json:"has_touch"`
	ScreenWidth       int      `json:"screenWidth,omitempty"`  //显示器宽度
	ScreenHeight      int      `json:"screenHeight,omitempty"` //显示器高度
	PositionX         int      `json:"positionX,omitempty"`    //浏览器在屏幕上的位置
	PositionY         int      `json:"positionY,omitempty"`    //浏览器在屏幕上的位置
}

type Screen struct {
	Width             int           `json:"width"`                  // 覆盖宽度，单位像素（最小0，最大10000000）。0表示禁用覆盖。
	Height            int           `json:"height"`                 // 覆盖高度，单位像素（最小0，最大10000000）。0表示禁用覆盖。
	DeviceScaleFactor float64       `json:"deviceScaleFactor"`      // 覆盖设备像素比。0表示禁用覆盖。
	Mobile            bool          `json:"mobile"`                 // 是否模拟移动设备，包括viewport meta标签、滚动条覆盖、文本自适应等。
	Scale             float64       `json:"scale,omitempty"`        // 应用于最终视图图像的缩放。实验性功能。
	ScreenWidth       int           `json:"screenWidth,omitempty"`  // 覆盖屏幕宽度，单位像素（最小0，最大10000000）。实验性功能。
	ScreenHeight      int           `json:"screenHeight,omitempty"` // 覆盖屏幕高度，单位像素（最小0，最大10000000）。实验性功能。
	PositionX         int           `json:"positionX,omitempty"`    // 覆盖视图在屏幕上的X位置，单位像素（最小0，最大10000000）。实验性功能。
	PositionY         int           `json:"positionY,omitempty"`    // 覆盖视图在屏幕上的Y位置，单位像素（最小0，最大10000000）。实验性功能。
	Viewport          *PageViewport `json:"viewport,omitempty"`     // 可见页面区域覆盖，页面不会感知此更改。实验性功能。
}

type PageViewport struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Scale  float64 `json:"scale"`
}

// 设置屏幕显示
func (obj *WebSock) EmulationSetScreenOverride(preCtx context.Context, device Screen) (RecvData, error) {
	deviceData, err := gson.Decode(device)
	if err != nil {
		return RecvData{}, err
	}
	return obj.send(preCtx, commend{
		Method: "Emulation.setDeviceMetricsOverride",
		Params: deviceData.RawMap(),
	})
}

// 设置屏幕显示
func (obj *WebSock) EmulationSetDeviceMetricsOverride(preCtx context.Context, device Device) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setDeviceMetricsOverride",
		Params: map[string]any{
			"width":             device.Viewport.Width,
			"height":            device.Viewport.Height,
			"deviceScaleFactor": device.DeviceScaleFactor,
			"mobile":            device.IsMobile,
			"has_touch":         device.HasTouch,
		},
	})
}

// 设置地理位置
func (obj *WebSock) EmulationSetGeolocationOverride(preCtx context.Context, latitude, longitude float64, accuracy int) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setGeolocationOverride",
		Params: map[string]any{
			"latitude":  latitude,
			"longitude": longitude,
			"accuracy":  accuracy, //float,100-2000。 定位精度 5	高精度 GPS，20	正常手机，100	WiFi 定位，1000	粗略 IP 定位
		},
	})
}

// 设置硬件并发数
func (obj *WebSock) EmulationSetHardwareConcurrencyOverride(preCtx context.Context, hardwareConcurrency int) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setHardwareConcurrencyOverride",
		Params: map[string]any{
			"hardwareConcurrency": hardwareConcurrency,
		},
	})
}

// 使用指定的区域设置覆盖默认主机系统区域设置。例如： en_US
func (obj *WebSock) EmulationSetLocaleOverride(preCtx context.Context, locale string) (RecvData, error) {
	key, _, ok := strings.Cut(locale, ",")
	if ok {
		locale = key
	}
	return obj.send(preCtx, commend{
		Method: "Emulation.setLocaleOverride",
		Params: map[string]any{
			"locale": strings.ReplaceAll(locale, "-", "_"),
		},
	})
}

// 使用指定的时区覆盖默认主机系统时区。
func (obj *WebSock) EmulationSetTimezoneOverride(preCtx context.Context, timezoneId string) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setTimezoneOverride",
		Params: map[string]any{
			"timezoneId": timezoneId,
		},
	})
}

// 设置页面空闲状态
func (obj *WebSock) EmulationSetActive(preCtx context.Context) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setIdleOverride",
		Params: map[string]any{
			"isUserActive":     true,  //用户是否活动
			"isScreenUnlocked": false, //屏幕是否上锁
		},
	})
}

// 设置cpu频率
func (obj *WebSock) EmulationSetCPUThrottlingRate(preCtx context.Context, rate float64) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setCPUThrottlingRate",
		Params: map[string]any{
			"rate": rate, //cpu 频率降低倍数，1 不降低，2，降低2倍
		},
	})
}

// 设置字体缩放比例
func (obj *WebSock) EmulationSetEmulatedOSTextScale(preCtx context.Context, scale float64) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setEmulatedOSTextScale",
		Params: map[string]any{
			"scale": scale, //字体缩放比例 0.9-1.1
		},
	})
}

// 处于焦点并激活页面
func (obj *WebSock) EmulationSetFocusEmulationEnabled(preCtx context.Context) (RecvData, error) {
	return obj.send(preCtx, commend{
		Method: "Emulation.setFocusEmulationEnabled",
		Params: map[string]any{
			"enabled": true,
		},
	})
}
