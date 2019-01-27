//
//  Medicine.swift
//  NWHacks
//
//  Created by Felipe on 2019-01-26.
//  Copyright © 2019 Felipe. All rights reserved.
//

import Foundation

struct Medicine: Codable {
    var Name: String
    var ID: String
    var _Type: String
    var Appereance: [String]
    var SideEffects: [String]
    private enum CodingKeys: String, CodingKey {
        case Name = "name"
        case ID = "med_id"
        case _Type = "type"
        case Appereance = "appereance"
        case SideEffects = "side_effects"
    }
}
