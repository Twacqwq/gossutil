package config

import "gossutil/pkg/config"

func init() {
	config.Add("aliyun", func() map[string]interface{} {
		return map[string]interface{}{
			"endpoint":          config.Env("ALIYUN_ENDPOINT", ""),
			"access_key_id":     config.Env("ALIYUN_ACCESS_KEY_ID", ""),
			"access_key_secret": config.Env("ALIYUN_ACCESS_KEY_SECRET", ""),
			"bucket":            config.Env("ALIYUN_BUCKET", ""),
			"upload_path":       config.Env("ALIYUN_UPLOAD_PATH", "public"),
		}
	})
}
