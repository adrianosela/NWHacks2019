//
//  VideoCell.swift
//  BeginnerTableView
//
//  Created by Sean Allen on 5/19/17.
//  Copyright © 2017 Sean Allen. All rights reserved.
//

import UIKit

class VideoCell: UITableViewCell {

    @IBOutlet weak var videoImageView: UIImageView!
    @IBOutlet weak var videoTitleLabel: UILabel!
    @IBOutlet weak var dosage: UILabel!
    @IBOutlet weak var doctor: UILabel!
    @IBOutlet weak var remaining: UILabel!
    
    
    func setPrescription(video: Video) {
        videoImageView.image = video.image
        videoTitleLabel.text = video.title
    }
    
}
