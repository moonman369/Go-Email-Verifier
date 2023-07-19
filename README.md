## This is a Go Based Email Verifier Tool

## API Documentation

### 1. VERIFY DOMAIN
- Base URL: `https://go-email-verifier.cleverapps.io/`
- Endpoint: `/verify`
- Method: `POST`
- Example Request Body:
  ```
  {
    "path": "doma.in",
  }
  ```
- Example Successful Response:
  ```
  200 OK
  {
    "domain": "doma.in",
    "hasMX": true,
    "hasSPF": true,
    "spfRecord": "spf_record",
    "hasDMARC": true,
    "dmarcRecord": "dmarc_record"
  }
  ```
<br>
