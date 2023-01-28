#!/bin/bash
docker compose up -d # stworzenie koneternu z bazą danych postgress
cd api && go run main.go&  # wystartowanie API
cd web && npm install && npm start& # wystartowanie apki webowej
curl -X POST localhost:8080/api/v1/data:csv-upload -F dataset=probne -F file=@./data/real_data.csv # zploadowanie danych prawdziwych
curl -X POST localhost:8080/api/v1/data:csv-upload -F dataset=predykaty -F file=@./data/pred_data.csv # zuploadowanie danych spredykowanych 
echo "otwórz przeglądarke i ciesz się panelem admina z danymi :D" #essa