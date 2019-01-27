//
//  NetworkController.swift
//  NWHacks
//
//  Created by Felipe on 2019-01-26.
//  Copyright © 2019 Felipe. All rights reserved.
//

import Foundation
import Alamofire

protocol NetworkControllerDelegate: class {
//    Notify app about server responses 
}

class NetworkController {
    private let baseURL = "http://localhost:8080"
    private let jsonDecoder = JSONDecoder()
    private let jsonEncoder = JSONEncoder()
}

// Mark: Doctors endpoints
extension NetworkController {
    
}

// Mark: Patients endpoints
extension NetworkController {
    
}

// Mark: Prescription endpoints
extension NetworkController {
    
}

