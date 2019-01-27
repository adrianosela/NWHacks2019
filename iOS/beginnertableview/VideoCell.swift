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
    
    
    func setPrescription(data: Tableview1) {
        videoImageView.image = #imageLiteral(resourceName: "traingularpill_green")
        videoTitleLabel.text = data.name
        dosage.text = data.dose
        doctor.text = data.doctor
        remaining.text = String(data.remaining!)
        
    }
    
}
