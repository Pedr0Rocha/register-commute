# Register Commute

CLI tool to help track commute days and make declaring costs easier

### Run

`go run cmd/main.go`

### Storage

It will generate a `.json` file in this format:

```json
[
 {
  "id": "8fb745bc-850f-4e4d-bca6-1cc9b30db4d2",
  "date": "2023-09-28",
  "transport": "Public Transport",
  "created_at": "2023-09-26 20:33:44.125353 +0000 UTC"
 },
 {
  "id": "55ab9589-897c-412e-a197-71724ed111fe",
  "date": "2023-09-27",
  "transport": "Bike",
  "created_at": "2023-09-26 20:33:39.988289 +0000 UTC"
 }
]
```

### Main Menu

![image](/images/main-menu.png)

### Diplay commutes

![image](/images/display-commutes.png)

### Register commute

![image](/images/register-commute.png)
