# gomud

## Usage

1. Build the software with `go build .`
2. Prepare the `.env` (you can copy from `.env.example` and change to match production values)
3. Execute it on CLI

## Available Commands

- student import: made to quickly import an institution and its students into legado.portalmud.com.br.
  **usage**:
  ```bash
  gomud student import --institution="{name of institution}" [--plan_id=2 --expiration_date="2026-05-31" --password="dacomud2025"]
  ```
  | argument | required | description | default |
  | --- | --- | --- | --- |
  | institution | yes | the name of the institution to be created. it's also the name of the CSV file to import students ({NAME}.csv) | - |
  | plan_id | no | the ID of the plan to apply to the students | 2 |
  | expiration_date | no | the final date when the students' plan will be valid | now + 1 year |
  | password | no | the password to be applied to the students account | {dacomud} + current YEAR |