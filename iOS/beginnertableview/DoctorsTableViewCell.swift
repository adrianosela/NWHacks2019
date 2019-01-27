//
//  VideoCell.swift
//  BeginnerTableView
//
//  Created by Sean Allen on 5/19/17.
//  Copyright © 2017 Sean Allen. All rights reserved.
//

import UIKit

class DoctorCell: UITableViewCell {
    
    @IBOutlet weak var videoImageView: UIImageView!
    @IBOutlet weak var videoTitleLabel: UILabel!
    @IBOutlet weak var profession: UILabel!

    var phone: String!
    
    @IBOutlet weak var callAction: UIButton!
    @IBAction func messageAction(_ sender: Any) {
    }
    
    func setPrescription(video: Video) {
        let url = URL(string: "https://felipeaccount.blob.core.windows.net/nwh/Doctor1.png")
        videoImageView.kf.indicatorType = .activity
        videoImageView.kf.setImage(with: url)
        videoTitleLabel.text = video.title
    }
    
}
