//
//  Patient.swift
//  NWHacks
//
//  Created by Felipe on 2019-01-26.
//  Copyright © 2019 Felipe. All rights reserved.
//

import Foundation

struct Patient: Codable {
    var ID: String
    var Name: String
    var Email: String
    var Phone: String
    var Prescriptions: [String]

    private enum CodingKeys: String, CodingKey {
        case ID = "patient_id"
        case Name = "name"
        case Email = "email"
        case Phone = "phone"
        case Prescriptions = "prescriptions"
    }
}
