# bSocial

Tehnički zadatak

## Disclaimer

Tri required stvari nisu napravljene u tasku

- Notification service kao poseban servis. Razlog: ne vidim smisao da bude odvojen servis, sa svojom svojom bazom, i da exposa 1 rest rutu. Jedina razlika bi bio to da consumerGrupa za notification service bude razlicita od consumerGrupe za telemetry service. ostalo sve bi bio prilicno Copy-paste iz ostalih servisa.

- ElasticSearch Query za telemetry servis. Napisao sam jedan query u telemetry/elastic_queries.txt. Ostale nisam znao jer prvi put koristim Elastic i jednostavno bi mi oduzelo previse vremena da naucim to iskljucivo za tehnicki zadatak.

- Dockerfile za servise, nikad nisam pakiro servis u docker. i isto kao i prije, previse bi vremena oduzelo

Napominjem ovo iz razloga sto pise da se treba sve uraditi, tako da eventualno ne morate ni gledati kod.

A u slucaju da je okej sto fali par stvari. idemo dalje

## Requirements

- go1.18 ili više(koristio sam generikse na jednom dijelu)
- docker-compose u slučaju da želite na taj način pokrenuti (mysql, kafka, elasticsearch), ako imate lokalno docker-compose nije potreban
- docker (isto nije potreban ako imate sve servise lokalno)

## How to run

svaka aplikacija zahtjeva config file u rootu servisa sa imenom `config.json`, ostavit cu u repozitoriju svoj config file tako da ne morate vlastiti pisati.

from repository folder run

```
docker-compose up
```

Pokretanje rest apia (exposan na portu :5000)

```
cd backend
go get
go run main.go
```

Pokretanje telemetry servisa

```
cd telemetry
go get
go run main.go
```

Također u root foderu repozitorija je fajl `bSocial.postman_collection.json`, to mozete importati u postman te cete imati mockupe svih mogućih ruta.

## Neke dizajn odluke

Procitao sam blog "The Clean Architecture" pa sam tio probati na taj nacin strukturirati kod, i cini mi se dosta okej (vjerovatno sam dosta odstupao, but hey...).

Koristio sam GORM (golang ORM) za mysql, nisam ga odavno koristio i tio sam se malo podsjetiti, ali naravno, znao bi RAW SQL querie pisati.

JWT je koristen za auth, posto nikada nisam implementirao na backendu, tio sam i to probati.

Imao sam u planu napraviti frontend, ali mi je ispalo dosta repetitivno, dosadno, i previse time-consuming (ako bi to napravio kako treba), tako da sam odustao od te ideje.

Postoji validacija requestova koja je built in u fiberu, tako da se validira request body i vrati error response ako nesto nije kako treba.

Mozda neki djelovi koda nisu consistent
