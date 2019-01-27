//
//  Doctors.swift
//  BeginnerTableView
//
//  Created by Sean Allen on 5/19/17.
//  Copyright © 2017 Sean Allen. All rights reserved.
//

import UIKit

class Doctors: UIViewController {
    
    @IBOutlet weak var tableView: UITableView!
    
    var videos: [Video] = []
    
    
    override func viewDidLoad() {
        super.viewDidLoad()
        videos = createArray()
    }
    
    
    func createArray() -> [Video] {
        
        let video1 = Video(image: #imageLiteral(resourceName: "Doctor2"), title: "Dr. Will Smith")
        let video2 = Video(image: #imageLiteral(resourceName: "Doctor3"), title: "Dr. Haze Cush")
        let video3 = Video(image: #imageLiteral(resourceName: "Doctor2"), title: "Dr. House")
        let video4 = Video(image: #imageLiteral(resourceName: "Doctor1"), title: "Google HAZE")
        let video5 = Video(image: #imageLiteral(resourceName: "Doctor3"), title: "Pill 5")
        let video6 = Video(image: #imageLiteral(resourceName: "Doctor3"), title: "Pill 6")
        
        return [video1, video2, video3, video4, video5, video6]
    }
}


extension Doctors: UITableViewDataSource, UITableViewDelegate {
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return videos.count
    }
    
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let video = videos[indexPath.row]
        let cell = tableView.dequeueReusableCell(withIdentifier: "DoctorCell") as! DoctorCell
        cell.setPrescription(video: video)
        
        return cell
    }
}





