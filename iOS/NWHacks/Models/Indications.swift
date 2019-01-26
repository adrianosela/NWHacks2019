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
    var Time: [String]
    
    private enum CodingKeys: String, CodingKey {
        case DaysPerWeek = "days_per_week"
        case TimesPerDay = "times_per_day"
        case Time = "times_of_day"
    }
}
