#!/bin/bash

# THIS SCRIPT IS RAN TO POPULATE A PATIENT WITH PRESCRIPTIONS AND DOCTORS

HOST="http://ezpillzz.azurewebsites.net"
DOCTOR_ID="d3a9e293-a3ff-4802-81c0-aac224f92834"
PATIENT_ID="3c1d77b9-79ff-4b09-a553-386664a634ba"

# create 6 prescriptions
rx_id_1=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Tylenol":{"dpw": 7, "tpd": 3,"tod" : [600, 1200, 1800]}},"amounts": {"Tylenol":90}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_2=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Ibuprofen":{"dpw": 7, "tpd": 3,"tod" : [600, 1200, 1800]}},"amounts": {"Ibuprofen":90}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_3=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Morphine":{"dpw": 5, "tpd": 1,"tod" : [600]}},"amounts": {"Morphine":45}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_4=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Codeine":{"dpw": 3, "tpd": 2,"tod" : [600, 1200]}},"amounts": {"Codeine":30}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_5=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Epinephrine":{"dpw": 3, "tpd": 2,"tod" : [600, 1200]}},"amounts": {"Epinephrine":30}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_6=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Heroine":{"dpw": 5, "tpd": 1,"tod" : [600]}},"amounts": {"Heroine":45}}' | http POST ${HOST}/prescription | jq -r .id)

#rx_id_1=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Contraceptive":{"dpw": 7, "tpd": 3,"tod" : [600, 1200, 1800]}},"amounts": {"Contraceptive":90}}' | http POST ${HOST}/prescription | jq -r .id)
#rx_id_2=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Xanax":{"dpw": 7, "tpd": 3,"tod" : [600, 1200, 1800]}},"amounts": {"Xanax":90}}' | http POST ${HOST}/prescription | jq -r .id)
#rx_id_3=$(echo '{"doctor": "'"$DOCTOR_ID"'","medicines": {"Medical Cannabis":{"dpw": 7, "tpd": 3,"tod" : [600, 1200, 1800]}},"amounts": {"Medicinal Cannabis":90}}' | http POST ${HOST}/prescription | jq -r .id)

# claim all 6 for patient
#echo '{"prescription_id": "'"$rx_id_1"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
#echo '{"prescription_id": "'"$rx_id_2"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
#echo '{"prescription_id": "'"$rx_id_3"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
#echo '{"prescription_id": "'"$rx_id_4"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
#echo '{"prescription_id": "'"$rx_id_5"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
#echo '{"prescription_id": "'"$rx_id_6"'","patient_id": "'"$PATIENT_ID"'"}' | http POST ${HOST}/claim
