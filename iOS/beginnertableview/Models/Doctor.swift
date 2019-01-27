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
    var OfficePhone: String?
    var Specializaion: String
    var PhotoURL: String?
    var Patients: [String]
    
    private enum CodingKeys: String, CodingKey {
        case ID = "id"
        case Name = "name"
        case Office = "office"
        case OfficePhone = "office_phone"
        case Specializaion = "specialization"
        case PhotoURL = "photo_url"
        case Patients = "patients"
    }
}

struct ListDoctors: Codable {
    var Doctors: [Doctor]
    
    private enum CodingKeys: String, CodingKey {
        case Doctors = "doctors"
    }
}
