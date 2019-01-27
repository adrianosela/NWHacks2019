#!/bin/bash

# THIS SCRIPT IS RAN TO POPULATE A PATIENT WITH PRESCRIPTIONS AND DOCTORS

HOST="http://sinful.azurewebsites.net"
DOCTOR_ID="9c0a6dab-287a-421a-84fd-cbe404021af8"
PATIENT_ID="4c0b926b-6976-4d6e-a9ca-3211375822b4"

# create 6 prescriptions
rx_id_1=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Tylenol":{"dpw": 7, "tpd": 3,"tod" : [600, 1200, 1800]}},"amounts": {"Tylenol":90}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_2=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Tylenol":{"dpw": 7, "tpd": 3,"tod" : [600, 1200, 1800]}},"amounts": {"Tylenol":90}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_3=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Morphine":{"dpw": 5, "tpd": 1,"tod" : [600]}},"amounts": {"Morphine":45}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_4=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Insulin":{"dpw": 3, "tpd": 2,"tod" : [600, 1200]}},"amounts": {"Insulin":30}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_5=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Insulin":{"dpw": 3, "tpd": 2,"tod" : [600, 1200]}},"amounts": {"Insulin":30}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_6=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Morphine":{"dpw": 5, "tpd": 1,"tod" : [600]}},"amounts": {"Morphine":45}}' | http POST ${HOST}/prescription | jq -r .id)

# claim all 6 for patient
echo '{"prescription_id": "'"$rx_id_1"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
echo '{"prescription_id": "'"$rx_id_2"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
echo '{"prescription_id": "'"$rx_id_3"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
echo '{"prescription_id": "'"$rx_id_4"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
echo '{"prescription_id": "'"$rx_id_5"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
echo '{"prescription_id": "'"$rx_id_6"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
