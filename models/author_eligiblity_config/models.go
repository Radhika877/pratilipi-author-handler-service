package authoreligiblityconfig

type AuthorEligiblityConfig struct {
	MinimumNoOfFollowers int `json:"MinimumNoOfFollowers"`
	MinimumNoOfPosts     int `json:"MinimumNoOfPosts"`
	WindowInDays         int `json:"WindowInDays"`
}

func (authorEligiblityConfig *AuthorEligiblityConfig) UpdateAuthorEligiblityConfigParams() map[string]interface{} {
	return map[string]interface{}{
		"minimumNoOfFollowers": authorEligiblityConfig.MinimumNoOfFollowers,
		"minimumNoOfPosts":     authorEligiblityConfig.MinimumNoOfPosts,
		"windowInDays":         authorEligiblityConfig.WindowInDays,
	}
}
