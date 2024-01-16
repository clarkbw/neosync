from buf.validate import validate_pb2 as _validate_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetSystemTransformersRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetSystemTransformersResponse(_message.Message):
    __slots__ = ("transformers",)
    TRANSFORMERS_FIELD_NUMBER: _ClassVar[int]
    transformers: _containers.RepeatedCompositeFieldContainer[SystemTransformer]
    def __init__(self, transformers: _Optional[_Iterable[_Union[SystemTransformer, _Mapping]]] = ...) -> None: ...

class GetUserDefinedTransformersRequest(_message.Message):
    __slots__ = ("account_id",)
    ACCOUNT_ID_FIELD_NUMBER: _ClassVar[int]
    account_id: str
    def __init__(self, account_id: _Optional[str] = ...) -> None: ...

class GetUserDefinedTransformersResponse(_message.Message):
    __slots__ = ("transformers",)
    TRANSFORMERS_FIELD_NUMBER: _ClassVar[int]
    transformers: _containers.RepeatedCompositeFieldContainer[UserDefinedTransformer]
    def __init__(self, transformers: _Optional[_Iterable[_Union[UserDefinedTransformer, _Mapping]]] = ...) -> None: ...

class GetUserDefinedTransformerByIdRequest(_message.Message):
    __slots__ = ("transformer_id",)
    TRANSFORMER_ID_FIELD_NUMBER: _ClassVar[int]
    transformer_id: str
    def __init__(self, transformer_id: _Optional[str] = ...) -> None: ...

class GetUserDefinedTransformerByIdResponse(_message.Message):
    __slots__ = ("transformer",)
    TRANSFORMER_FIELD_NUMBER: _ClassVar[int]
    transformer: UserDefinedTransformer
    def __init__(self, transformer: _Optional[_Union[UserDefinedTransformer, _Mapping]] = ...) -> None: ...

class CreateUserDefinedTransformerRequest(_message.Message):
    __slots__ = ("account_id", "name", "description", "type", "source", "transformer_config")
    ACCOUNT_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    TRANSFORMER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    account_id: str
    name: str
    description: str
    type: str
    source: str
    transformer_config: TransformerConfig
    def __init__(self, account_id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., type: _Optional[str] = ..., source: _Optional[str] = ..., transformer_config: _Optional[_Union[TransformerConfig, _Mapping]] = ...) -> None: ...

class CreateUserDefinedTransformerResponse(_message.Message):
    __slots__ = ("transformer",)
    TRANSFORMER_FIELD_NUMBER: _ClassVar[int]
    transformer: UserDefinedTransformer
    def __init__(self, transformer: _Optional[_Union[UserDefinedTransformer, _Mapping]] = ...) -> None: ...

class DeleteUserDefinedTransformerRequest(_message.Message):
    __slots__ = ("transformer_id",)
    TRANSFORMER_ID_FIELD_NUMBER: _ClassVar[int]
    transformer_id: str
    def __init__(self, transformer_id: _Optional[str] = ...) -> None: ...

class DeleteUserDefinedTransformerResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class UpdateUserDefinedTransformerRequest(_message.Message):
    __slots__ = ("transformer_id", "name", "description", "transformer_config")
    TRANSFORMER_ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    TRANSFORMER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    transformer_id: str
    name: str
    description: str
    transformer_config: TransformerConfig
    def __init__(self, transformer_id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., transformer_config: _Optional[_Union[TransformerConfig, _Mapping]] = ...) -> None: ...

class UpdateUserDefinedTransformerResponse(_message.Message):
    __slots__ = ("transformer",)
    TRANSFORMER_FIELD_NUMBER: _ClassVar[int]
    transformer: UserDefinedTransformer
    def __init__(self, transformer: _Optional[_Union[UserDefinedTransformer, _Mapping]] = ...) -> None: ...

class IsTransformerNameAvailableRequest(_message.Message):
    __slots__ = ("account_id", "transformer_name")
    ACCOUNT_ID_FIELD_NUMBER: _ClassVar[int]
    TRANSFORMER_NAME_FIELD_NUMBER: _ClassVar[int]
    account_id: str
    transformer_name: str
    def __init__(self, account_id: _Optional[str] = ..., transformer_name: _Optional[str] = ...) -> None: ...

class IsTransformerNameAvailableResponse(_message.Message):
    __slots__ = ("is_available",)
    IS_AVAILABLE_FIELD_NUMBER: _ClassVar[int]
    is_available: bool
    def __init__(self, is_available: bool = ...) -> None: ...

class UserDefinedTransformer(_message.Message):
    __slots__ = ("id", "name", "description", "data_type", "source", "config", "created_at", "updated_at", "account_id")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DATA_TYPE_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    CONFIG_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    ACCOUNT_ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    description: str
    data_type: str
    source: str
    config: TransformerConfig
    created_at: _timestamp_pb2.Timestamp
    updated_at: _timestamp_pb2.Timestamp
    account_id: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., description: _Optional[str] = ..., data_type: _Optional[str] = ..., source: _Optional[str] = ..., config: _Optional[_Union[TransformerConfig, _Mapping]] = ..., created_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., updated_at: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ..., account_id: _Optional[str] = ...) -> None: ...

class SystemTransformer(_message.Message):
    __slots__ = ("name", "description", "data_type", "source", "config")
    NAME_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    DATA_TYPE_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    CONFIG_FIELD_NUMBER: _ClassVar[int]
    name: str
    description: str
    data_type: str
    source: str
    config: TransformerConfig
    def __init__(self, name: _Optional[str] = ..., description: _Optional[str] = ..., data_type: _Optional[str] = ..., source: _Optional[str] = ..., config: _Optional[_Union[TransformerConfig, _Mapping]] = ...) -> None: ...

class TransformerConfig(_message.Message):
    __slots__ = ("generate_email_config", "transform_email_config", "generate_bool_config", "generate_card_number_config", "generate_city_config", "generate_e164_phone_number_config", "generate_first_name_config", "generate_float64_config", "generate_full_address_config", "generate_full_name_config", "generate_gender_config", "generate_int64_phone_number_config", "generate_int64_config", "generate_last_name_config", "generate_sha256hash_config", "generate_ssn_config", "generate_state_config", "generate_street_address_config", "generate_string_phone_number_config", "generate_string_config", "generate_unixtimestamp_config", "generate_username_config", "generate_utctimestamp_config", "generate_uuid_config", "generate_zipcode_config", "transform_e164_phone_number_config", "transform_first_name_config", "transform_float64_config", "transform_full_name_config", "transform_int64_phone_number_config", "transform_int64_config", "transform_last_name_config", "transform_phone_number_config", "transform_string_config", "passthrough_config", "nullconfig", "user_defined_transformer_config", "generate_default_config", "transform_javascript_config", "generate_categorical_config")
    GENERATE_EMAIL_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_EMAIL_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_BOOL_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_CARD_NUMBER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_CITY_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_E164_PHONE_NUMBER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_FIRST_NAME_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_FLOAT64_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_FULL_ADDRESS_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_FULL_NAME_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_GENDER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_INT64_PHONE_NUMBER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_INT64_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_LAST_NAME_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_SHA256HASH_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_SSN_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_STATE_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_STREET_ADDRESS_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_STRING_PHONE_NUMBER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_STRING_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_UNIXTIMESTAMP_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_USERNAME_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_UTCTIMESTAMP_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_UUID_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_ZIPCODE_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_E164_PHONE_NUMBER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_FIRST_NAME_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_FLOAT64_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_FULL_NAME_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_INT64_PHONE_NUMBER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_INT64_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_LAST_NAME_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_PHONE_NUMBER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_STRING_CONFIG_FIELD_NUMBER: _ClassVar[int]
    PASSTHROUGH_CONFIG_FIELD_NUMBER: _ClassVar[int]
    NULLCONFIG_FIELD_NUMBER: _ClassVar[int]
    USER_DEFINED_TRANSFORMER_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_DEFAULT_CONFIG_FIELD_NUMBER: _ClassVar[int]
    TRANSFORM_JAVASCRIPT_CONFIG_FIELD_NUMBER: _ClassVar[int]
    GENERATE_CATEGORICAL_CONFIG_FIELD_NUMBER: _ClassVar[int]
    generate_email_config: GenerateEmail
    transform_email_config: TransformEmail
    generate_bool_config: GenerateBool
    generate_card_number_config: GenerateCardNumber
    generate_city_config: GenerateCity
    generate_e164_phone_number_config: GenerateE164PhoneNumber
    generate_first_name_config: GenerateFirstName
    generate_float64_config: GenerateFloat64
    generate_full_address_config: GenerateFullAddress
    generate_full_name_config: GenerateFullName
    generate_gender_config: GenerateGender
    generate_int64_phone_number_config: GenerateInt64PhoneNumber
    generate_int64_config: GenerateInt64
    generate_last_name_config: GenerateLastName
    generate_sha256hash_config: GenerateSha256Hash
    generate_ssn_config: GenerateSSN
    generate_state_config: GenerateState
    generate_street_address_config: GenerateStreetAddress
    generate_string_phone_number_config: GenerateStringPhoneNumber
    generate_string_config: GenerateString
    generate_unixtimestamp_config: GenerateUnixTimestamp
    generate_username_config: GenerateUsername
    generate_utctimestamp_config: GenerateUtcTimestamp
    generate_uuid_config: GenerateUuid
    generate_zipcode_config: GenerateZipcode
    transform_e164_phone_number_config: TransformE164PhoneNumber
    transform_first_name_config: TransformFirstName
    transform_float64_config: TransformFloat64
    transform_full_name_config: TransformFullName
    transform_int64_phone_number_config: TransformInt64PhoneNumber
    transform_int64_config: TransformInt64
    transform_last_name_config: TransformLastName
    transform_phone_number_config: TransformPhoneNumber
    transform_string_config: TransformString
    passthrough_config: Passthrough
    nullconfig: Null
    user_defined_transformer_config: UserDefinedTransformerConfig
    generate_default_config: GenerateDefault
    transform_javascript_config: TransformJavascript
    generate_categorical_config: GenerateCategorical
    def __init__(self, generate_email_config: _Optional[_Union[GenerateEmail, _Mapping]] = ..., transform_email_config: _Optional[_Union[TransformEmail, _Mapping]] = ..., generate_bool_config: _Optional[_Union[GenerateBool, _Mapping]] = ..., generate_card_number_config: _Optional[_Union[GenerateCardNumber, _Mapping]] = ..., generate_city_config: _Optional[_Union[GenerateCity, _Mapping]] = ..., generate_e164_phone_number_config: _Optional[_Union[GenerateE164PhoneNumber, _Mapping]] = ..., generate_first_name_config: _Optional[_Union[GenerateFirstName, _Mapping]] = ..., generate_float64_config: _Optional[_Union[GenerateFloat64, _Mapping]] = ..., generate_full_address_config: _Optional[_Union[GenerateFullAddress, _Mapping]] = ..., generate_full_name_config: _Optional[_Union[GenerateFullName, _Mapping]] = ..., generate_gender_config: _Optional[_Union[GenerateGender, _Mapping]] = ..., generate_int64_phone_number_config: _Optional[_Union[GenerateInt64PhoneNumber, _Mapping]] = ..., generate_int64_config: _Optional[_Union[GenerateInt64, _Mapping]] = ..., generate_last_name_config: _Optional[_Union[GenerateLastName, _Mapping]] = ..., generate_sha256hash_config: _Optional[_Union[GenerateSha256Hash, _Mapping]] = ..., generate_ssn_config: _Optional[_Union[GenerateSSN, _Mapping]] = ..., generate_state_config: _Optional[_Union[GenerateState, _Mapping]] = ..., generate_street_address_config: _Optional[_Union[GenerateStreetAddress, _Mapping]] = ..., generate_string_phone_number_config: _Optional[_Union[GenerateStringPhoneNumber, _Mapping]] = ..., generate_string_config: _Optional[_Union[GenerateString, _Mapping]] = ..., generate_unixtimestamp_config: _Optional[_Union[GenerateUnixTimestamp, _Mapping]] = ..., generate_username_config: _Optional[_Union[GenerateUsername, _Mapping]] = ..., generate_utctimestamp_config: _Optional[_Union[GenerateUtcTimestamp, _Mapping]] = ..., generate_uuid_config: _Optional[_Union[GenerateUuid, _Mapping]] = ..., generate_zipcode_config: _Optional[_Union[GenerateZipcode, _Mapping]] = ..., transform_e164_phone_number_config: _Optional[_Union[TransformE164PhoneNumber, _Mapping]] = ..., transform_first_name_config: _Optional[_Union[TransformFirstName, _Mapping]] = ..., transform_float64_config: _Optional[_Union[TransformFloat64, _Mapping]] = ..., transform_full_name_config: _Optional[_Union[TransformFullName, _Mapping]] = ..., transform_int64_phone_number_config: _Optional[_Union[TransformInt64PhoneNumber, _Mapping]] = ..., transform_int64_config: _Optional[_Union[TransformInt64, _Mapping]] = ..., transform_last_name_config: _Optional[_Union[TransformLastName, _Mapping]] = ..., transform_phone_number_config: _Optional[_Union[TransformPhoneNumber, _Mapping]] = ..., transform_string_config: _Optional[_Union[TransformString, _Mapping]] = ..., passthrough_config: _Optional[_Union[Passthrough, _Mapping]] = ..., nullconfig: _Optional[_Union[Null, _Mapping]] = ..., user_defined_transformer_config: _Optional[_Union[UserDefinedTransformerConfig, _Mapping]] = ..., generate_default_config: _Optional[_Union[GenerateDefault, _Mapping]] = ..., transform_javascript_config: _Optional[_Union[TransformJavascript, _Mapping]] = ..., generate_categorical_config: _Optional[_Union[GenerateCategorical, _Mapping]] = ...) -> None: ...

class GenerateEmail(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class TransformEmail(_message.Message):
    __slots__ = ("preserve_domain", "preserve_length")
    PRESERVE_DOMAIN_FIELD_NUMBER: _ClassVar[int]
    PRESERVE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    preserve_domain: bool
    preserve_length: bool
    def __init__(self, preserve_domain: bool = ..., preserve_length: bool = ...) -> None: ...

class GenerateBool(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateCardNumber(_message.Message):
    __slots__ = ("valid_luhn",)
    VALID_LUHN_FIELD_NUMBER: _ClassVar[int]
    valid_luhn: bool
    def __init__(self, valid_luhn: bool = ...) -> None: ...

class GenerateCity(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateDefault(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateE164PhoneNumber(_message.Message):
    __slots__ = ("min", "max")
    MIN_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    min: int
    max: int
    def __init__(self, min: _Optional[int] = ..., max: _Optional[int] = ...) -> None: ...

class GenerateFirstName(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateFloat64(_message.Message):
    __slots__ = ("randomize_sign", "min", "max", "precision")
    RANDOMIZE_SIGN_FIELD_NUMBER: _ClassVar[int]
    MIN_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    PRECISION_FIELD_NUMBER: _ClassVar[int]
    randomize_sign: bool
    min: float
    max: float
    precision: int
    def __init__(self, randomize_sign: bool = ..., min: _Optional[float] = ..., max: _Optional[float] = ..., precision: _Optional[int] = ...) -> None: ...

class GenerateFullAddress(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateFullName(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateGender(_message.Message):
    __slots__ = ("abbreviate",)
    ABBREVIATE_FIELD_NUMBER: _ClassVar[int]
    abbreviate: bool
    def __init__(self, abbreviate: bool = ...) -> None: ...

class GenerateInt64PhoneNumber(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateInt64(_message.Message):
    __slots__ = ("randomize_sign", "min", "max")
    RANDOMIZE_SIGN_FIELD_NUMBER: _ClassVar[int]
    MIN_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    randomize_sign: bool
    min: int
    max: int
    def __init__(self, randomize_sign: bool = ..., min: _Optional[int] = ..., max: _Optional[int] = ...) -> None: ...

class GenerateLastName(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateSha256Hash(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateSSN(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateState(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateStreetAddress(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateStringPhoneNumber(_message.Message):
    __slots__ = ("include_hyphens",)
    INCLUDE_HYPHENS_FIELD_NUMBER: _ClassVar[int]
    include_hyphens: bool
    def __init__(self, include_hyphens: bool = ...) -> None: ...

class GenerateString(_message.Message):
    __slots__ = ("min", "max")
    MIN_FIELD_NUMBER: _ClassVar[int]
    MAX_FIELD_NUMBER: _ClassVar[int]
    min: int
    max: int
    def __init__(self, min: _Optional[int] = ..., max: _Optional[int] = ...) -> None: ...

class GenerateUnixTimestamp(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateUsername(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateUtcTimestamp(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GenerateUuid(_message.Message):
    __slots__ = ("include_hyphens",)
    INCLUDE_HYPHENS_FIELD_NUMBER: _ClassVar[int]
    include_hyphens: bool
    def __init__(self, include_hyphens: bool = ...) -> None: ...

class GenerateZipcode(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class TransformE164PhoneNumber(_message.Message):
    __slots__ = ("preserve_length",)
    PRESERVE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    preserve_length: bool
    def __init__(self, preserve_length: bool = ...) -> None: ...

class TransformFirstName(_message.Message):
    __slots__ = ("preserve_length",)
    PRESERVE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    preserve_length: bool
    def __init__(self, preserve_length: bool = ...) -> None: ...

class TransformFloat64(_message.Message):
    __slots__ = ("randomization_range_min", "randomization_range_max")
    RANDOMIZATION_RANGE_MIN_FIELD_NUMBER: _ClassVar[int]
    RANDOMIZATION_RANGE_MAX_FIELD_NUMBER: _ClassVar[int]
    randomization_range_min: float
    randomization_range_max: float
    def __init__(self, randomization_range_min: _Optional[float] = ..., randomization_range_max: _Optional[float] = ...) -> None: ...

class TransformFullName(_message.Message):
    __slots__ = ("preserve_length",)
    PRESERVE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    preserve_length: bool
    def __init__(self, preserve_length: bool = ...) -> None: ...

class TransformInt64PhoneNumber(_message.Message):
    __slots__ = ("preserve_length",)
    PRESERVE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    preserve_length: bool
    def __init__(self, preserve_length: bool = ...) -> None: ...

class TransformInt64(_message.Message):
    __slots__ = ("randomization_range_min", "randomization_range_max")
    RANDOMIZATION_RANGE_MIN_FIELD_NUMBER: _ClassVar[int]
    RANDOMIZATION_RANGE_MAX_FIELD_NUMBER: _ClassVar[int]
    randomization_range_min: int
    randomization_range_max: int
    def __init__(self, randomization_range_min: _Optional[int] = ..., randomization_range_max: _Optional[int] = ...) -> None: ...

class TransformLastName(_message.Message):
    __slots__ = ("preserve_length",)
    PRESERVE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    preserve_length: bool
    def __init__(self, preserve_length: bool = ...) -> None: ...

class TransformPhoneNumber(_message.Message):
    __slots__ = ("preserve_length", "include_hyphens")
    PRESERVE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    INCLUDE_HYPHENS_FIELD_NUMBER: _ClassVar[int]
    preserve_length: bool
    include_hyphens: bool
    def __init__(self, preserve_length: bool = ..., include_hyphens: bool = ...) -> None: ...

class TransformString(_message.Message):
    __slots__ = ("preserve_length",)
    PRESERVE_LENGTH_FIELD_NUMBER: _ClassVar[int]
    preserve_length: bool
    def __init__(self, preserve_length: bool = ...) -> None: ...

class Passthrough(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class Null(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class TransformJavascript(_message.Message):
    __slots__ = ("code",)
    CODE_FIELD_NUMBER: _ClassVar[int]
    code: str
    def __init__(self, code: _Optional[str] = ...) -> None: ...

class UserDefinedTransformerConfig(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class ValidateUserJavascriptCodeRequest(_message.Message):
    __slots__ = ("account_id", "code")
    ACCOUNT_ID_FIELD_NUMBER: _ClassVar[int]
    CODE_FIELD_NUMBER: _ClassVar[int]
    account_id: str
    code: str
    def __init__(self, account_id: _Optional[str] = ..., code: _Optional[str] = ...) -> None: ...

class ValidateUserJavascriptCodeResponse(_message.Message):
    __slots__ = ("valid",)
    VALID_FIELD_NUMBER: _ClassVar[int]
    valid: bool
    def __init__(self, valid: bool = ...) -> None: ...

class GenerateCategorical(_message.Message):
    __slots__ = ("categories",)
    CATEGORIES_FIELD_NUMBER: _ClassVar[int]
    categories: str
    def __init__(self, categories: _Optional[str] = ...) -> None: ...