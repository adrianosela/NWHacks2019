//
//  Doctors.swift
//  BeginnerTableView
//
//  Created by Sean Allen on 5/19/17.
//  Copyright © 2017 Sean Allen. All rights reserved.
//

import UIKit
import Alamofire

class Doctors: UIViewController {
    
    @IBOutlet weak var tableView: UITableView!
    private let jsonDecoder = JSONDecoder()
    private let jsonEncoder = JSONEncoder()
    var data: ListDoctors?
    
    
    override func viewDidLoad() {
        super.viewDidLoad()
        let id = UserDefaults.standard.string(forKey: "id")

        getDoctors(userId: id!)
    }
    
    func getDoctors(userId: String) {
        let parameters: [String: Any] = [
            "id": userId
        ]
        Alamofire.request("http://ezpillzz.azurewebsites.net"+"/patient_doctors/" + userId, method: .get, parameters: parameters).responseJSON {
            response in
            print(response)
            if let json = response.data, let listDoctors = try? self.jsonDecoder.decode(ListDoctors.self, from: json) {
                print(listDoctors, listDoctors.Doctors.count)
                self.data = listDoctors
                self.tableView.reloadData()
            } else {
                print(response.error ?? "Unknow Error")
            }
        }
    }
    
}


extension Doctors: UITableViewDataSource, UITableViewDelegate {
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return data?.Doctors.count ?? 0
    }
    
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let doc = data?.Doctors[indexPath.row]
        let cell = tableView.dequeueReusableCell(withIdentifier: "DoctorCell") as! DoctorCell
        if((doc) != nil) {cell.setPrescription(doc: doc!)}
        
        return cell
    }
}





