//
//  VideoListScreen.swift
//  BeginnerTableView
//
//  Created by Sean Allen on 5/19/17.
//  Copyright © 2017 Sean Allen. All rights reserved.
//

import UIKit

class VideoListScreen: UIViewController {
    
    @IBOutlet weak var tableView: UITableView!
    
    var videos: [Video] = []

    
    override func viewDidLoad() {
        super.viewDidLoad()
        videos = createArray()
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
}


extension VideoListScreen: UITableViewDataSource, UITableViewDelegate {
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return videos.count
    }
    
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let video = videos[indexPath.row]
        let cell = tableView.dequeueReusableCell(withIdentifier: "VideoCell") as! VideoCell
        cell.setVideo(video: video)
        
        return cell
    }
}





