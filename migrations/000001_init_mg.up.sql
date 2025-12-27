CREATE TABLE certificate_tls (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    url TEXT NOT NULL UNIQUE,
    version INTEGER NOT NULL,
    
    not_before TIMESTAMPTZ NOT NULL,
    not_after TIMESTAMPTZ NOT NULL,

    c_subject TEXT NOT NULL,
    c_issuer TEXT NOT NULL,
    serial_number TEXT NOT NULL
);