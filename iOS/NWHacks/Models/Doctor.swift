//
//  Doctor.swift
//  NWHacks
//
//  Created by Felipe on 2019-01-26.
//  Copyright © 2019 Felipe. All rights reserved.
//

import Foundation

struct Doctor: Codable {
    var ID: String
    var Name: String
    var Office: String
    var Specializaion: String
    var Patients: [String]
    
    private enum CodingKeys: String, CodingKey {
        case ID = "doctor_id"
        case Name = "name"
        case Office = "office"
        case Specializaion = "specialization"
        case Patients = "patients"
    }
}
