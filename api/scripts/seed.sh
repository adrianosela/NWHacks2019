#!/bin/bash

# THIS SCRIPT IS RAN TO POPULATE THE DB - ENOUGH FOR A DEMO

HOST="http://ezpills.azurewebsites.net"

# create 5 doctors
dr_id_1=$(echo '{"name": "Felipe Ballesteros Md.","office": "2265 West 16th Avenue, Vancouver BC, V6R4F9","specialization": "cardiology"}' | http POST ${HOST}/doctor | jq -r .id)
dr_id_2=$(echo '{"name": "House Md.","office": "4527 West 10th Avenue, Vancouver BC, V6D3G9","specialization": "neurology"}' | http POST ${HOST}/doctor | jq -r .id)
dr_id_3=$(echo '{"name": "Collin Horricks Md.", "office": "2345 Birney Ave, Vancouver BC, V6K2J9", "specialization": "dermatology"}' | http POST ${HOST}/doctor | jq -r .id)
dr_id_4=$(echo '{"name": "Jenny Riley Md.","office": "927 Granville St, Vancouver BC, V2D35F","specialization": "infectious diseases"}' | http POST ${HOST}/doctor | jq -r .id)
dr_id_5=$(echo '{"name": "Farshid Agarebparast Md.", "office": "3647 Davie St, Vancouver BC, V5Z3E9", "addiction psychiatry": ""}' | http POST ${HOST}/doctor | jq -r .id)

# create 3 prescriptions
rx_id_1=$(echo '{"doctor": "'"$dr_id_1"'","medicines": {"Tylenol":{"dpw": 7, "tpd": 3,"tod" : [600, 1200, 1800]}},"amounts": {"Tylenol":90}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_2=$(echo '{"doctor": "'"$dr_id_2"'","medicines": {"Morphine":{"dpw": 5, "tpd": 1,"tod" : [600]}},"amounts": {"Morphine":45}}' | http POST ${HOST}/prescription | jq -r .id)
rx_id_3=$(echo '{"doctor": "'"$dr_id_3"'","medicines": {"Insulin":{"dpw": 3, "tpd": 2,"tod" : [600, 1200]}},"amounts": {"Insulin":30}}' | http POST ${HOST}/prescription | jq -r .id)

# claim 2 with new patients
pt_id_1=$(echo '{"name": "Mike Tyson","email": "miketyson@gmail.com", "phone": "(778) 681 2638", "prescription_id": "'"$rx_id_1"'"}' | http POST ${HOST}/patient | jq -r .id)
pt_id_2=$(echo '{"name": "John Appleseed","email": "john.appleseed@gmail.com", "phone": "(604) 681 5825", "prescription_id": "'"$rx_id_2"'"}' | http POST ${HOST}/patient | jq -r .id)
