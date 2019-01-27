//
//  Tables.swift
//  BeginnerTableView
//
//  Created by Felipe on 2019-01-27.
//  Copyright © 2019 Sean Allen. All rights reserved.
//

import Foundation

struct Tableview1 {
    var name: String?
    var dose: String?
    var doctor: String?
    var doctorId: String?
    var remaining: Int?
    
    init(name: String, dose: String, doctorId: String, remaining: Int) {
        self.name = name
        self.dose = dose
        self.doctorId = doctorId
        self.doctor = nil
        self.remaining = remaining
    }
}
