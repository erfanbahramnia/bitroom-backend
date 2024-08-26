package constants

import "time"

const CsrfCookieName = "csrf_cookie"
const Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const Numbers = "0123456789"
const OtpLength = 5
const CacheItemTimeExpiration = 3 * time.Minute
const CachePurgesTimeExpiratoin = 5 * time.Minute
