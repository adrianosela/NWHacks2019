//
//  Prescription.swift
//  NWHacks
//
//  Created by Felipe on 2019-01-26.
//  Copyright © 2019 Felipe. All rights reserved.
//

import Foundation

struct Prescription: Codable {
    var ID: String
    var AddedAt: Int
    var Medicines: String
    var Remaining: String
    var Claimed: Bool
    var Patient: String
    private enum CodingKeys: String, CodingKey {
        case ID = "rx_id"
        case AddedAt = "dosage"
        case Medicines = "meds"
        case Remaining = "remaining"
        case Claimed = "claimed"
        case Patient = "patient"
    }
}
