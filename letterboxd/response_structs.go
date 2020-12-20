package letterboxd

// AccessToken holds information to sign every request with a valid access token
// data provided by the getAccessToken() function
type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

// MemberAccount holds data for a Letterboxd member
// http://api-docs.letterboxd.com/#/definitions/MemberAccount
type MemberAccount struct {
	EmailAddress                             string        `json:"emailAddress"`
	EmailAddressValidated                    bool          `json:"emailAddressValidated"`
	PrivateAccount                           bool          `json:"privateAccount"`
	IncludeInPeopleSection                   bool          `json:"includeInPeopleSection"`
	PrivateWatchlist                         bool          `json:"privateWatchlist"`
	EmailWhenFollowed                        bool          `json:"emailWhenFollowed"`
	EmailAvailability                        bool          `json:"emailAvailability"`
	EmailBuyAvailability                     bool          `json:"emailBuyAvailability"`
	EmailRentAvailability                    bool          `json:"emailRentAvailability"`
	EmailComments                            bool          `json:"emailComments"`
	EmailNews                                bool          `json:"emailNews"`
	EmailRushes                              bool          `json:"emailRushes"`
	DevicesRegisteredForPushNotifications    []string      `json:"devicesRegisteredForPushNotifications"`
	PushNotificationsForGeneralAnnouncements bool          `json:"pushNotificationsForGeneralAnnouncements"`
	PushNotificationsForComments             bool          `json:"pushNotificationsForComments"`
	PushNotificationsForListLikes            bool          `json:"pushNotificationsForListLikes"`
	PushNotificationsForReviewLikes          bool          `json:"pushNotificationsForReviewLikes"`
	PushNotificationsForNewFollowers         bool          `json:"pushNotificationsForNewFollowers"`
	PushNotificationsForAvailability         bool          `json:"pushNotificationsForAvailability"`
	PushNotificationsForBuyAvailability      bool          `json:"pushNotificationsForBuyAvailability"`
	PushNotificationsForRentAvailability     bool          `json:"pushNotificationsForRentAvailability"`
	CanComment                               bool          `json:"canComment"`
	Suspended                                bool          `json:"suspended"`
	CanCloneLists                            bool          `json:"canCloneLists"`
	CanFilterActivity                        bool          `json:"canFilterActivity"`
	AuthorizedSharingServicesForLists        []interface{} `json:"authorizedSharingServicesForLists"`
	AuthorizedSharingServicesForReviews      []interface{} `json:"authorizedSharingServicesForReviews"`
	Member                                   struct {
		ID          string `json:"id"`
		Username    string `json:"username"`
		DisplayName string `json:"displayName"`
		ShortName   string `json:"shortName"`
		Pronoun     struct {
			ID                  string `json:"id"`
			Label               string `json:"label"`
			SubjectPronoun      string `json:"subjectPronoun"`
			ObjectPronoun       string `json:"objectPronoun"`
			PossessiveAdjective string `json:"possessiveAdjective"`
			PossessivePronoun   string `json:"possessivePronoun"`
			Reflexive           string `json:"reflexive"`
		} `json:"pronoun"`
		Avatar struct {
			Sizes []struct {
				Width  int    `json:"width"`
				Height int    `json:"height"`
				URL    string `json:"url"`
			} `json:"sizes"`
		} `json:"avatar"`
		CommentPolicy    string        `json:"commentPolicy"`
		AccountStatus    string        `json:"accountStatus"`
		MemberStatus     string        `json:"memberStatus"`
		HideAdsInContent bool          `json:"hideAdsInContent"`
		HideAds          bool          `json:"hideAds"`
		BioLbml          string        `json:"bioLbml"`
		FavoriteFilms    []interface{} `json:"favoriteFilms"`
		Links            []struct {
			Type string `json:"type"`
			ID   string `json:"id"`
			URL  string `json:"url"`
		} `json:"links"`
		PrivateWatchlist bool   `json:"privateWatchlist"`
		Bio              string `json:"bio"`
	} `json:"member"`
}
