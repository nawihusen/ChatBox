package domain

// import (
// 	"be-service-saksi-management/saksi/repository/grpc/realcountmgt"
// 	"context"
// 	"time"
// )

// type Saksi struct {
// 	ID           int64
// 	MemberNo     string
// 	RecruiterID  int64
// 	ProvCode     int64
// 	KabCode      int64
// 	KecCode      int64
// 	KelCode      int64
// 	TpsIDRequest int64
// 	TpsID        *int64
// 	Status       string `json:"status" form:"status" validate:"required"`
// 	DtmCrt       time.Time
// 	DtmUpd       time.Time
// }

// type SaksiRequest struct {
// 	MemberNo string `json:"member_no" form:"member_no" validate:"required"`
// 	TpsNo    int64  `json:"tps_no" form:"tps_no" validate:"required"`
// }

// // SaksiCandidate model
// type SaksiCandidate struct {
// 	ID          int32
// 	MemberNo    string
// 	PhoneNumber string
// 	QRCode      string
// 	DtmCrt      time.Time
// 	DtmUpd      time.Time
// }

// // AccountResponseDTO is saksi response DTO
// type AccountResponseDTO struct {
// 	ID              int32           `json:"id"`
// 	MemberNo        string          `json:"member_no"`
// 	NIK             string          `json:"nik"`
// 	FullName        string          `json:"full_name"`
// 	PhoneNumber     string          `json:"phone_number"`
// 	Address         string          `json:"address"`
// 	DPRTCoordinator DPRTCoordinator `json:"dprt_coordinator"`
// 	SelfieImage     string          `json:"selfie_image"`
// 	KTPImage        string          `json:"ktp_image"`
// }

// // AuthResponseDTO is auth response DTO
// type AuthResponseDTO struct {
// 	MemberNo    string `json:"member_no"`
// 	NIK         string `json:"nik"`
// 	FullName    string `json:"full_name"`
// 	Role        string `json:"role"`
// 	MemberImage string `json:"member_image"`
// 	Kecamatan   string `json:"kecamatan"`
// 	Kelurahan   string `json:"kelurahan"`
// 	TpsNo       *int64 `json:"tps_no,omitempty"`
// 	Attendance  *bool  `json:"attendance,omitempty"`
// 	Token       string `json:"token"`
// }

// // MemberOnline is saksi response from member online service via gRPC
// type Session struct {
// 	UserID          int64  `json:"user_id" redis:"user_id"`
// 	RecruiterRole   bool   `json:"recruiter_role" redis:"recruiter_role"`
// 	SaksiRole       bool   `json:"saksi_role" redis:"saksi_role"`
// 	CoordinatorRole bool   `json:"coordinator_role" redis:"coordinator_role"`
// 	RecruiterID     int64  `json:"recruiter_id" redis:"recruiter_id"`
// 	SaksiID         int64  `json:"saksi_id" redis:"saksi_id"`
// 	CoordinatorID   int64  `json:"coordinator_id" redis:"coordinator_id"`
// 	MemberNo        string `json:"member_no" redis:"member_no"`
// 	NIK             string `json:"nik" redis:"nik"`
// 	FullName        string `json:"full_name" redis:"full_name"`
// 	BirthDate       string `json:"birth_date" redis:"birth_date"`
// 	BirthPlace      string `json:"birth_place" redis:"birth_place"`
// 	Gender          string `json:"gender" redis:"gender"`
// 	Religion        string `json:"religion" redis:"religion"`
// 	MarriedStatus   string `json:"married_status" redis:"married_status"`
// 	FullAddress     string `json:"full_address" redis:"full_address"`
// 	ProvID          int64  `json:"prov_id" redis:"prov_id"`
// 	KabID           int64  `json:"kab_id" redis:"kab_id"`
// 	KecID           int64  `json:"kec_id" redis:"kec_id"`
// 	KelID           int64  `json:"kel_id" redis:"kel_id"`
// 	Email           string `json:"email" redis:"email"`
// 	PhoneNumber     string `json:"phone_number" redis:"phone_number"`
// 	TpsNo           int64  `json:"tps_no" redis:"tps_no"`
// 	Attendance      bool   `json:"attendance" redis:"attendance"`
// }

// // DPRTCoordinator model
// type DPRTCoordinator struct {
// 	MemberNo string `json:"member_no"`
// 	Name     string `json:"name"`
// }

// // SaksiDPRT model
// type SaksiDPRT struct {
// 	ID          int64
// 	MemberNo    string
// 	PhoneNumber string
// 	Role        string
// 	ProvCode    int64
// 	KabCode     int64
// 	KecCode     int64
// 	KelCode     int64
// 	TpsID       *int64
// 	DtmCrt      time.Time
// 	DtmUpd      time.Time
// }

// // MemberOnline is saksi response from member online service via gRPC
// type MemberOnline struct {
// 	UserID        int64  `json:"user_id"`
// 	MemberNo      string `json:"member_no"`
// 	NIK           string `json:"nik"`
// 	FullName      string `json:"full_name"`
// 	BirthDate     string `json:"birth_date"`
// 	BirthPlace    string `json:"birth_place"`
// 	Gender        string `json:"gender"`
// 	Religion      string `json:"religion"`
// 	MarriedStatus string `json:"married_status"`
// 	FullAddress   string `json:"full_address"`
// 	ProvID        int64  `json:"prov_id"`
// 	KabID         int64  `json:"kab_id"`
// 	KecID         int64  `json:"kec_id"`
// 	KelID         int64  `json:"kel_id"`
// 	Email         string `json:"email"`
// }

// // SaksiValidationResponseDTO is SaksiValidationResponseDTO response
// type SaksiValidationResponseDTO struct {
// 	Nik            string          `json:"nik"`
// 	FullName       string          `json:"full_name"`
// 	PhoneNumber    string          `json:"phone_number"`
// 	Address        string          `json:"address"`
// 	DPRTCordinator DPRTCoordinator `json:"dprt_coordinator"`
// 	QRQode         string          `json:"qr_code"`
// }

// // Authorization model
// type Authorization struct {
// 	UserID      int64  `json:"user_id"`
// 	PhoneNumber string `json:"phone_number"`
// }

// type User struct {
// 	PhoneNumber string `json:"phone_number"`
// }

// // TPS model
// type TPS struct {
// 	ID       int64
// 	ProvCode int64
// 	KabCode  int64
// 	KecCode  int64
// 	KelCode  int64
// 	TpsNo    int64
// 	DtmCrt   time.Time
// 	DtmUpd   time.Time
// }

// // SaksiDPRTRequestDTO is saksi DPRT request DTO
// type SaksiDPRTRequestDTO struct {
// 	CandidateID int64 `json:"candidate_id" form:"candidate_id"`
// 	TpsNo       int64 `json:"tps_no" form:"tps_no"`
// }

// // CreateSaksiResponseDTO is CreateSaksi response
// type CreateSaksiResponseDTO struct {
// 	ID       int64  `json:"id"`
// 	MemberNo string `json:"member_no"`
// 	TpsNo    int64  `json:"tps_no"`
// }

// // ListSaksiDTO is List Saksi
// type ListSaksiDTO struct {
// 	NIK         string `json:"nik"`
// 	MemberNo    string `json:"member_no"`
// 	PhoneNumber string `json:"phone_number"`
// 	FullName    string `json:"full_name"`
// 	TpsNo       int64  `json:"tps_no"`
// 	Address     string `json:"address"`
// 	MemberImage string `json:"member_image"`
// 	SelfieImage string `json:"selfie_image"`
// 	KTPImage    string `json:"ktp_image"`
// }

// // SaksiDTO is Saksi DTO
// type SaksiDTO struct {
// 	ID          int64  `json:"id"`
// 	MemberNo    string `json:"member_no"`
// 	PhoneNumber string `json:"phone_number"`
// 	FullName    string `json:"full_name"`
// 	TpsNo       int64  `json:"tps_no"`
// 	MemberImage string `json:"member_image"`
// 	SelfieImage string `json:"selfie_image"`
// 	KTPImage    string `json:"ktp_image"`
// }

// type TpsRequested struct {
// 	ID    int64 `json:"id"`
// 	TpsNo int64 `json:"tps_no"`
// }

// type TpsAssigned struct {
// 	ID    *int64 `json:"id"`
// 	TpsNo *int64 `json:"tps_no"`
// }
// type GetSaksiRecruiter struct {
// 	ID          int64  `json:"id"`
// 	NIK         string `json:"nik"`
// 	PhoneNumber string `json:"phone_number"`
// 	FullName    string `json:"full_name"`
// }
// type GetSaksiRespon struct {
// 	ID           int64             `json:"id"`
// 	MemberNo     string            `json:"member_no"`
// 	PhoneNumber  string            `json:"phone_number"`
// 	FullName     string            `json:"full_name"`
// 	Status       string            `json:"status"`
// 	NIK          string            `json:"nik"`
// 	Recruiter    GetSaksiRecruiter `json:"recruiter"`
// 	Provinsi     WilayahDTO        `json:"provinsi"`
// 	Kabupaten    WilayahDTO        `json:"kabupaten"`
// 	Kecamatan    WilayahDTO        `json:"kecamatan"`
// 	Kelurahan    WilayahDTO        `json:"kelurahan"`
// 	TpsRequested TpsRequested      `json:"tps_requested"`
// 	TpsAssigned  TpsAssigned       `json:"tps_assigned"`
// }

// type SelectSaksiRespon struct {
// 	ID             int64
// 	MemberNo       string
// 	RecruiterID    int64
// 	ProvCode       int64
// 	KabCode        int64
// 	KecCode        int64
// 	KelCode        int64
// 	TpsIdRequested int64
// 	TpsId          int64
// 	Status         string
// }

// // AreaByCodeRequest model
// type AreaByCodeRequest struct {
// 	ProvCode   int32 `json:"prov_code" form:"prov_code"`
// 	KabCode    int32 `json:"kab_code" form:"kab_code"`
// 	KecCode    int32 `json:"kec_code" form:"kec_code"`
// 	KelCode    int32 `json:"kel_code" form:"kel_code"`
// 	PostalCode bool  `json:"postal_code" form:"postal_code"`
// 	Legacy     bool  `json:"legacy" form:"legacy"`
// 	Location   bool  `json:"location" form:"location"`
// }

// // AreaResponse model
// type AreaResponse struct {
// 	Provinsi    Area  `json:"provinsi"`
// 	Kabupaten   Area  `json:"kabupaten"`
// 	Kecamatan   Area  `json:"kecamatan"`
// 	Kelurahan   Area  `json:"kelurahan"`
// 	Level       int32 `json:"level"`
// 	CombineCode string
// 	CombineName string
// 	PostalCode  string
// 	Latitude    float32
// 	Longitude   float32
// }

// // Area model
// type Area struct {
// 	Code     int32  `json:"code"`
// 	Name     string `json:"name"`
// 	LegacyId int32  `json:"legacy_id"`
// }

// // GetAreaByCodeLegacyRequest model
// type GetAreaByCodeLegacyRequest struct {
// 	ProvId     int32 `json:"prov_id" form:"prov_id"`
// 	KabId      int32 `json:"kab_id" form:"kab_id"`
// 	KecId      int32 `json:"kec_id" form:"kec_id"`
// 	KelId      int32 `json:"kel_id" form:"kel_id"`
// 	PostalCode bool  `json:"postal_code" form:"postal_code"`
// 	Location   bool  `json:"location" form:"location"`
// }

// // MetaData model
// type MetaData struct {
// 	TotalData uint   `json:"total_data"`
// 	TotalPage uint   `json:"total_page"`
// 	Page      uint   `json:"page"`
// 	Limit     uint   `json:"limit"`
// 	Sort      string `json:"sort"`
// 	Order     string `json:"order"`
// }

// type SupporterData struct {
// 	UserID      int64
// 	FullName    string
// 	PhoneNumber string
// }

// type SaksiAdditionalData struct {
// 	RecruiterID     int64 `json:"recruiter_id" redis:"recruiter_id"`
// 	SaksiID         int64 `json:"saksi_id" redis:"saksi_id"`
// 	CoordinatorID   int64 `json:"coordinator_id" redis:"coordinator_id"`
// 	RoleRecruiter   bool  `json:"recruiter_role" redis:"recruiter_role"`
// 	RoleSaksi       bool  `json:"saksi_role" redis:"saksi_role"`
// 	RoleCoordinator bool  `json:"coordinator_role" redis:"coordinator_role"`
// 	TpsNo           int64 `redis:"tps_no"`
// 	Attendance      bool  `redis:"attendance"`
// }

// type DPRTDTO struct {
// 	ID          int64     `json:"id"`
// 	MemberNo    string    `json:"member_no"`
// 	PhoneNumber string    `json:"phone_number"`
// 	NIK         string    `json:"nik"`
// 	FullName    string    `json:"full_name"`
// 	MemberImage string    `json:"member_image"`
// 	Provinsi    Provinsi  `json:"provinsi"`
// 	Kabupaten   Kabupaten `json:"kabupaten"`
// 	Kecamatan   Kecamatan `json:"kecamatan"`
// 	Kelurahan   Kelurahan `json:"kelurahan"`
// }

// type Response struct {
// 	MetaData MetaData  `json:"metadata"`
// 	Data     []DPRTDTO `json:"data"`
// }

// type MemberAuthorization struct {
// 	UserID  int64 `json:"user_id"`
// 	IsAdmin bool  `json:"is_admin"`
// }

// type GetAllSaksiRequest struct {
// 	Page        int64  `json:"page"`
// 	Limit       int64  `json:"limit"`
// 	ProvCode    int64  `json:"prov_code"`
// 	KabCode     int64  `json:"kab_code"`
// 	KecCode     int64  `json:"kec_code"`
// 	KelCode     int64  `json:"kel_code"`
// 	RecruiterId int64  `json:"recruiter_id"`
// 	Status      string `json:"status"`
// 	Sort        string `json:"sort"`
// 	Order       string `json:"order"`
// }

// type GetAllSaksiResponse struct {
// 	MetaData MetaData         `json:"metadata"`
// 	Data     []GetAllSaksiDTO `json:"data"`
// }

// type GetAllSaksiDTO struct {
// 	ID           int64        `json:"id"`
// 	MemberNo     string       `json:"member_no"`
// 	Name         string       `json:"name"`
// 	Recruiter    RecruiterDTO `json:"recruiter"`
// 	Provinsi     DetailArea   `json:"provinsi"`
// 	Kabupaten    DetailArea   `json:"kabupaten"`
// 	Kecamatan    DetailArea   `json:"kecamatan"`
// 	Kelurahan    DetailArea   `json:"kelurahan"`
// 	TpsRequested TpsRequested `json:"tps_requested"`
// 	TpsAssigned  *TpsAssigned `json:"tps_assigned,omitempty"`
// 	Status       string       `json:"status"`
// }

// type RecruiterDTO struct {
// 	ID   int64  `json:"id"`
// 	Name string `json:"name"`
// }

// type DetailArea struct {
// 	Code int64  `json:"code"`
// 	Name string `json:"name"`
// }

// type UpdateSaksiTpsRequest struct {
// 	TpsNo int64 `json:"tps_no" form:"tps_no" validate:"required"`
// }

// type Lookup struct {
// 	NIK         string    `json:"nik"`
// 	FullName    string    `json:"full_name"`
// 	PhoneNumber string    `json:"phone_number"`
// 	Provinsi    Provinsi  `json:"provinsi"`
// 	Kabupaten   Kabupaten `json:"kabupaten"`
// 	Kecamatan   Kecamatan `json:"kecamatan"`
// 	Kelurahan   Kelurahan `json:"kelurahan"`
// }

// type ResponseTokenB2BMotionPay struct {
// 	AccessToken string `json:"accessToken"`
// }

// type MotionPayRequest struct {
// 	Token   string  `json:"token"`
// 	Session Session `json:"session"`
// }
// type UpdateStatusRequest struct {
// 	ReferelUrl string `json:"referer_url" form:"referer_url"`
// }

// // AccountUsecase is Account usecase
// type AccountUsecase interface {
// 	Authorize(ctx context.Context, token string) (authorization Authorization, member MemberOnline, add SaksiAdditionalData, err error)
// 	AuthorizeAuth(ctx context.Context, token string) (authorization Authorization, err error)
// 	GetValidation(ctx context.Context, qrCode *string) (saksi SaksiCandidate, candidate MemberOnline, coordinator MemberOnline, err error)
// 	DeleteSaksiValidation(ctx context.Context, id int64) error
// 	CreateSaksi(ctx context.Context, candidateID int64, tpsNo int64) (lastID int64, memberNo string, err error)
// 	ListSaksiDPRT(ctx context.Context, session *Session) ([]ListSaksiDTO, error)
// 	SumSaksi(ctx context.Context, prov, kab, kec, kel int64) (sum int64, err error)
// 	GetAllSaksi(ctx context.Context, limit, page int64, provCode, kabCode, kecCode, kelCode, sort, order string) ([]SaksiDTO, int64, error)
// 	GetSaksi(ctx context.Context, id int64, idRec int64, token string) (respon GetSaksiRespon, err error)
// 	InsertDPRT(ctx context.Context, dprt *SaksiDPRT) (err error)
// 	UpdateDPRT(ctx context.Context, dprt *SaksiDPRT) (err error)
// 	DeleteDPRT(ctx context.Context, no string) (err error)
// 	GetListDPRT(ctx context.Context, meta *MetaData, page, limit, prov, kab, kec, kel *int64, sort, order *string) (dprt []DPRTDTO, err error)
// 	PostSaksi(ctx context.Context, saksi *SaksiRequest, reqruiterID int64) (err error)
// 	PutSaksiStatus(ctx context.Context, saksi *Saksi) (err error)
// 	GetSemuaSaksi(ctx context.Context, req GetAllSaksiRequest) (res GetAllSaksiResponse, err error)
// 	GetSaksiRecruiter(ctx context.Context, req GetAllSaksiRequest, recruiterID int64) (res GetAllSaksiResponse, err error)
// 	UpdateSaksiTps(ctx context.Context, id, tpsNo int64) (err error)
// 	LookUpMember(ctx context.Context, memberNo string) (member Lookup, err error)
// 	PostAuthMotionPay(ctx context.Context, session MotionPayRequest) (response ResponseMotionPay, err error)
// 	EditSaksiStatus(ctx context.Context, req UpdateStatusRequest, token string, session Session) (err error)
// }

// // AccountMySQLRepository is Account repository in MySQL
// type AccountMySQLRepository interface {
// 	GetOneValidation(ctx context.Context, qrCode *string) (SaksiCandidate, error)
// 	GetSaksiDPRT(ctx context.Context, provID int64, kabID int64, kecID int64, kelID int64) (SaksiDPRT, error)
// 	InsertCandidate(ctc context.Context, candidate SaksiCandidate) (err error)
// 	SelectCoordinator(ctc context.Context, provID, KabID, KecID, KelID int64) (member_no string, err error)
// 	RemoveSaksiValidation(ctx context.Context, id int64) error
// 	// InsertSaksi(ctx context.Context, memberNo, phoneNumber, role string, provID, kabID, kecID, kelID, tpsID int64) (int64, error)
// 	GetCandidateByID(ctx context.Context, candidateID int64) (SaksiCandidate, error)
// 	GetTPSByCode(ctx context.Context, provCode, kabCode, kecCode, kelCode, tpsNo int64) (TPS, error)
// 	GetAllSaksiDPRT(ctx context.Context, provCode, kabCode, kecCode, kelCode int64) ([]ListSaksiDTO, error)
// 	GetSaksiDPRTByMemberNo(ctx context.Context, memberNo string) (SaksiDPRT, error)
// 	SumSaksi(ctx context.Context, prov, kab, kec, kel int64) (sum int64, err error)
// 	GetSaksiDPRTByTPS(ctx context.Context, tpsID int64) (SaksiDPRT, error)
// 	SelectAllSaksi(ctx context.Context, limit, page int64, provCode, kabCode, kecCode, kelCode, sort, order string) ([]SaksiDTO, error)
// 	SelectSaksi(ctx context.Context, id int64, recruiterID int64) (SelectSaksiRespon, error)
// 	CountAllSaksi(ctx context.Context, provCode, kabCode, kecCode, kelCode string) (res int64, err error)
// 	CountSaksiDPRTByMemberNo(ctx context.Context, memberNo string) (found int64, err error)
// 	InsertDPRT(ctx context.Context, dprt *SaksiDPRT) (err error)
// 	GetDPRTByArea(ctx context.Context, provCode, kabCode, kecCode, kelCode int64) (err error)
// 	UpdateDPRT(ctx context.Context, dprt *SaksiDPRT) (err error)
// 	DeleteDPRT(ctx context.Context, no string) (err error)
// 	SelectAllDPRT(ctx context.Context, limit, offset, prov, kab, kec, kel *int64, sort, order *string) (dprt []SaksiDPRT, err error)
// 	SumDPRT(ctx context.Context, prov, kab, kec, kel *int64) (total uint, err error)
// 	InsertSaksi(ctx context.Context, saksi *Saksi) (err error)
// 	UpdateSaksiStatus(ctx context.Context, saksi *Saksi) (err error)
// 	IsExistSaksi(ctx context.Context, memberNo string) (err error)
// 	GetSaksiByMemberNo(ctx context.Context, memberNo string) (saksi Saksi, err error)
// 	UserRequestQouta(ctx context.Context, memberNo string) (err error)
// 	TpsRequestQouta(ctx context.Context, tpsID int64) (err error)
// 	SelectSemuaSaksi(ctx context.Context, req GetAllSaksiRequest) (saksi []Saksi, err error)
// 	CountSemuaSaksi(ctx context.Context, req GetAllSaksiRequest) (count int64, err error)
// 	SelectSemuaSaksiRecruiter(ctx context.Context, req GetAllSaksiRequest, recruiterID int64) (saksi []Saksi, err error)
// 	CountSemuaSaksiRecruiter(ctx context.Context, req GetAllSaksiRequest, recruiterID int64) (count int64, err error)
// 	GetSaksiByID(ctx context.Context, id int64) (saksi Saksi, err error)
// 	EditSaksiTps(ctx context.Context, id, tpsNo int64) (err error)
// 	CheckSaksiExist(ctx context.Context, provCode, kabCode, kecCode, kelCode, tpsNo int64) (err error)
// 	CheckDoubleRequest(ctx context.Context, saksi *Saksi) (err error)
// 	SelectSaksiByMemberNo(ctx context.Context, memberNo string) (res SelectSaksiRespon, err error)
// 	UpdateStatusSaksi(ctx context.Context, memberNo string) (err error)
// }

// type AccountRedisRepository interface {
// 	SetSession(ctx context.Context, token string, add SaksiAdditionalData) (err error)
// 	GetSession(ctx context.Context, token string) (add SaksiAdditionalData, err error)
// 	DeleteSession(ctx context.Context, token string) (err error)
// 	AttendanceSession(ctx context.Context, token string) (err error)
// }

// // StructureGRPCRepository is Saksi repository in gRPC
// type SaksiGRPCRepository interface {
// 	GetDetailMemberByNo(ctx context.Context, memberNo string) (member MemberOnline, err error)
// 	GetDetailMemberByID(ctx context.Context, id int64) (member MemberOnline, err error)
// 	GetSupporterData(ctx context.Context, member MemberOnline) (data SupporterData, err error)
// }

// // SaksiGRPCAuthRepository is Structure repository in gRPC
// type SaksiGRPCAuthRepository interface {
// 	GetFromAuth(ctc context.Context, token string) (auth Authorization, err error)
// }

// type UserGRPCUserRepository interface {
// 	GetFromUser(ctx context.Context, userId int64) (user User, err error)
// }

// // AreaGRPCRepository is Area repository in gRPC
// type AreaGRPCRepository interface {
// 	GetSubAreaByCodeLegacy(ctx context.Context, in *GetAreaByCodeLegacyRequest) (area AreaResponse, err error)
// 	GetAreaByCode(ctx context.Context, areaReq *AreaByCodeRequest) (area AreaResponse, err error)
// 	ExchangeIDToCode(ctx context.Context, provID, KabID, KecID, KelID int64) (provCode, KabCode, KecCode, KelCode int64, err error)
// 	GetAreaByCodeBulk(ctx context.Context, code *[]SaksiDPRT) (area []AreaResponse, err error)
// }

// // RealCountGRPCRepository is Real count mgt repository in gRPC
// type RealCountGRPCRepository interface {
// 	GetTPSStatus(ctx context.Context, provID, kabID, kecID, kelID int64) (res *realcountmgt.GetTPSStatusResponse, err error)
// }
