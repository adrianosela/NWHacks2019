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
    private let baseURL = "http://slimjim.azurewebsites.net"
    private let jsonDecoder = JSONDecoder()
    private let jsonEncoder = JSONEncoder()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        videos = createArray()
        getAllPrescription(userId: "ef04ca5b-def4-4dce-8b9c-bd3b487e2117")
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
                print(listPrescription)
                Alamofire.request(self.self.baseURL+path, method: .get, parameters: parameters).responseJSON {
                    response in
                    print(response)
                    if let json = response.data, let listPrescription = try? self.jsonDecoder.decode(ListPrescription.self, from: json) {
                        print(listPrescription)
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
        return videos.count
    }
    
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let video = videos[indexPath.row]
        let cell = tableView.dequeueReusableCell(withIdentifier: "VideoCell") as! VideoCell
        cell.setPrescription(video: video)
        
        return cell
    }
}





