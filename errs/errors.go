package errs

const (
	ErrReadRequest        = "Ошибка чтения запроса"
	ErrWriteResponse      = "Ошибка записи ответа"
	ErrGeneralAuthError   = "Ошибка авторизации запроса"
	ErrDatabaseRequest    = "Ошибка при запросе к БД"
	ErrBadRequest         = "Некорректный запрос"
	ErrAPIRequest         = "Ошибка при запросе к сторонней API"
	ErrNoPermission       = "Отсутствует разрешение"
	ErrJSONDecode         = "Ошибка при чтении JSON объекта"
	ErrJSONEncode         = "Ошибка при записи JSON объекта"
	ErrRequestError       = "Ошибка выполнения запроса"
	ErrInvalidRequest     = "Ошибка валидации запроса"
	ErrNegativeEmployeeID = "Некорректный EmployeeID"
	ErrUnsupported        = "Необработанная ошибка"
)

// auth errors
const (
	ErrNoAuthHeader        = "An authorization header is required"
	ErrIncorrectAuthHeader = "Incorrect authorization header"
	ErrInvalidToken        = "Invalid authorization token"
	ErrPubKeyNotFound      = "Public key not found"
	ErrParseToken          = "Failed to parse token"
	ErrWrongEmployeeID     = "Wrong employee ID"
)

// permission errors
const (
	ErrPermissionNotFound = "Required permission was not found"
)
