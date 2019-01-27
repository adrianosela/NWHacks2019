//
//  Prescription.swift
//  NWHacks
//
//  Created by Felipe on 2019-01-26.
//  Copyright © 2019 Felipe. All rights reserved.
//

import Foundation
import UIKit

struct Prescription: Codable {
    var ID: String
    var AddedAt: Int
    var Medicines: [String:Indications]
    var Remaining: [String:Int]
    var Claimed: Bool
    var Patient: String
    var Doctor: String
    
    private enum CodingKeys: String, CodingKey {
        case ID = "id"
        case AddedAt = "added_at"
        case Medicines = "medicines"
        case Remaining = "remaining"
        case Claimed = "claimed"
        case Patient = "patient"
        case Doctor = "doctor"
    }
}

struct ListPrescription: Codable {
    var Prescriptions: [Prescription]
    
    private enum CodingKeys: String, CodingKey {
        case Prescriptions = "prescriptions"
    }
}

