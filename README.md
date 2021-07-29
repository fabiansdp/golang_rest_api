# Golang REST API
> Golang REST API using Gin, Gorm and MySQL

## Daftar Isi
* [Kontributor](#kontributor)
* [Penjelasan](#penjelasan)
* [Routes](#routes)

## Kontributor
Fabian Savero Diaz Pranoto (13519140)

## Penjelasan
Sebuah REST API sebagai tugas seleksi Asisten Lab Programming. Digunakan untuk menghandle request sebuah toko dorayaki.

## Routes
### Dorayaki Routes
1. Get Dorayaki List (GET /api/dorayakis)
```
Example Response:
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": [
        {
            "id": 1,
            "rasa": "Coklat",
            "deskripsi": "Enak",
            "gambar": "1c4494bc-bb9d-4577-9cbc-09e064117ece.jpg"
        },
        {
            "id": 2,
            "rasa": "Stroberi",
            "deskripsi": "Manis",
            "gambar": "c49fbd2e-7e33-4108-b06c-581a99aaed69.JPG"
        }
    ]
}
```
2. Get Single Dorayaki (GET /api/dorayakis/:id)
```
Example Response:
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 1,
        "rasa": "Coklat",
        "deskripsi": "Enak",
        "gambar": "1c4494bc-bb9d-4577-9cbc-09e064117ece.jpg"
    }
}
```
3. Create Dorayaki (POST /api/dorayakis)
```
Form Data:
{
    "rasa": "Coklat",
    "deskripsi": "Enak Banget",
    "gambar": File
}
```
4. Update Dorayaki (PATCH /api/dorayakis/:id)
```
Form Data:
{
    "rasa": "Coklat",
    "deskripsi": "Enak Banget",
    "gambar": File
}
```
5. Delete Dorayaki (DELETE /api/dorayakis/:id)
```
Example Response:
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "id": 4,
        "rasa": "Fabian",
        "deskripsi": "Fabian",
        "gambar": "1c4494bc-bb9d-4577-9cbc-09e064117ece.jpg"
    }
}
```

### Shop Routes
1. Get Shop List (GET /api/shops)
```
Example Response:
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": [
        {
            "id": 1,
            "created_at": "2021-07-29T12:39:20.725Z",
            "updated_at": "2021-07-29T12:39:20.725Z",
            "nama": "Toko Fabian",
            "jalan": "Gayungsari",
            "kecamatan": "Gayungan",
            "provinsi": "Jawa Timur"
        }
    ]
}
```
2. Get Single Shop (GET /api/shops/:id)
```
Example Response:
{
    "status": true,
    "message": "OK",
    "errors": null,
    "data": {
        "shop_info": {
            "id": 1,
            "created_at": "2021-07-29T12:39:20.725Z",
            "updated_at": "2021-07-29T12:39:20.725Z",
            "nama": "Toko Fabian",
            "jalan": "Gayungsari",
            "kecamatan": "Gayungan ",
            "provinsi": "Jawa Timur"
        },
        "inventory": [
            {
                "id": 1,
                "rasa": "Coklat",
                "deskripsi": "Enak",
                "gambar": "1c4494bc-bb9d-4577-9cbc-09e064117ece.jpg",
                "quantity": 1000
            }
        ]
    }
}
```
3. Create Shop (POST /api/shops)
```
Request Payload Example:
{
    "nama": "Toko Fabian",
    "jalan": "Gayungsari",
    "kecamatan": "Gayungan",
    "provinsi": "Jawa Barat"
}
```
4. Update Shop (PATCH /api/shops/:id)
```
Request Payload Example:
{
    "nama": "Toko Fabian",
    "jalan": "Gayungsari",
    "kecamatan": "Gayungan",
    "provinsi": "Jawa Barat"
}
```
5. Delete Shop (DELETE /api/shops/:id)
```
Example Response:
{
    "status": true,
    "message": "Deleted",
    "errors": null,
    "data": {}
}
```
### Inventory Routes
1. Add Shop Inventory (POST /api/inventory)
```
Request Payload Example:
{
    "dorayaki_id": 2,
    "shop_id": 2,
    "quantity": 100
}
```
2. Edit Inventory Quantity (PUT /api/inventory/:shopId)
```
Request Payload Example:
{
    "dorayaki_id": 1,
    "shop_id": 1,
    "quantity": 90
}
```
3. Move Inventory To Other Shop (PATCH /api/inventory/:shopId)
```
Request Payload Example:
{
    "dorayaki_id": 5,
    "recipient_id": 1,
    "quantity": 10
}
```
4. Delete Inventory Dorayaki (DELETE /api/inventory/:shopId)
```
Example Response:
{
    "status": true,
    "message": "Deleted",
    "errors": null,
    "data": {}
}
```
### File Routes
1. Get Dorayaki Photo (GET /api/file/:filename)
```
Response berupa sebuah foto
```