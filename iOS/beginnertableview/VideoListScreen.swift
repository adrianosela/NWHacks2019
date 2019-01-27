//
//  VideoListScreen.swift
//  BeginnerTableView
//
//  Created by Sean Allen on 5/19/17.
//  Copyright © 2017 Sean Allen. All rights reserved.
//

import UIKit
import Alamofire

class VideoListScreen: UIViewController {
    
    @IBOutlet weak var tableView: UITableView!
    
    var videos: [Video] = []
    private let baseURL = "http://applepen.azurewebsites.net"
    private let jsonDecoder = JSONDecoder()
    private let jsonEncoder = JSONEncoder()
    var tdata = [Tableview1]()

    override func viewDidLoad() {
        super.viewDidLoad()
        videos = createArray()
        getAllPrescription(userId: "3c1d77b9-79ff-4b09-a553-386664a634ba")
    }
    
    
    func createArray() -> [Video] {
        
        let video1 = Video(image: #imageLiteral(resourceName: "traingularpill_green"), title: "Advil")
        let video2 = Video(image: #imageLiteral(resourceName: "pill_green"), title: "Morphine")
        let video3 = Video(image: #imageLiteral(resourceName: "circularpill_green"), title: "NWHacks KUSH")
        let video4 = Video(image: #imageLiteral(resourceName: "pill_green"), title: "Google HAZE")
        let video5 = Video(image: #imageLiteral(resourceName: "traingularpill_green"), title: "Pill 5")
        let video6 = Video(image: #imageLiteral(resourceName: "circularpill_green"), title: "Pill 6")
        return [video1, video2, video3, video4, video5, video6]
    }
    
    public func getAllPrescription(userId: String) {
        let path = "/patient_prescriptions/"+userId
        let parameters: [String: Any] = [
            "id": userId
        ]
        
        Alamofire.request(baseURL+path, method: .get, parameters: parameters).responseJSON {
            response in
            print(response)
            if let json = response.data, let listPrescription = try? self.jsonDecoder.decode(ListPrescription.self, from: json) {
                var info = [Tableview1]()
                for number in 0..<(listPrescription._Prescriptions.count) {
                    let dose = String(Int.random(in: 0 ..< 45))
                    let doctorId = listPrescription._Prescriptions[number].Doctor
                    let Remaining = listPrescription._Prescriptions[number].Remaining
                    let medicines = listPrescription._Prescriptions[number].Medicines
                    for (key, val) in medicines {
                        let name = key
                        let remaining = Remaining[name]
                        let t = Tableview1(name: name, dose: dose, doctorId: doctorId, remaining: remaining!)
                        info.append(t)
                    }
                }
                print(info)
                Alamofire.request("http://applepen.azurewebsites.net"+"/patient_doctors/"+userId, method: .get, parameters: parameters).responseJSON {
                    response in
                    print(response)
                    if let json = response.data, let listDoctors = try? self.jsonDecoder.decode(ListDoctors.self, from: json) {
                        print(listDoctors, listDoctors.Doctors.count)
                        
                        for (number2) in 0..<(listDoctors.Doctors.count) {
                            let id = listDoctors.Doctors[number2].ID
                            let name = listDoctors.Doctors[number2].Name
                            for number2 in 0..<(info.count){
                                if(id == info[number2].doctorId!) { info[number2].doctor = name}
                            }
                        }
                        self.tdata = info
                        self.tableView.reloadData()
                        print("AAAA", info)
                        
                    } else {
                        print(response.error ?? "Unknow Error")
                    }
                }
            } else {
                print(response.error ?? "Unknow Error")
            }
        }
    }
    
}


extension VideoListScreen: UITableViewDataSource, UITableViewDelegate {
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return tdata.count
    }
    
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let video = tdata[indexPath.row]
        let cell = tableView.dequeueReusableCell(withIdentifier: "VideoCell") as! VideoCell
        
        if(tdata.count > 0) {cell.setPrescription(data: tdata[indexPath.row])}
        
        return cell
    }
}





