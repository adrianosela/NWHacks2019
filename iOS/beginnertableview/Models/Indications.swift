//
//  Indications.switf.swift
//  NWHacks
//
//  Created by Felipe on 2019-01-26.
//  Copyright © 2019 Felipe. All rights reserved.
//

import Foundation

struct Indications: Codable {
    var DaysPerWeek: Int
    var TimesPerDay: Int
    var Time: [Int]
    
    private enum CodingKeys: String, CodingKey {
        case DaysPerWeek = "dpw"
        case TimesPerDay = "tpd"
        case Time = "tod"
    }
}
