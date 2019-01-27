//
//  logoutViewController.swift
//  BeginnerTableView
//
//  Created by Felipe on 2019-01-27.
//  Copyright © 2019 Sean Allen. All rights reserved.
//

import UIKit
import Kingfisher


class logoutViewController: UIViewController {
    
    @IBOutlet weak var videoImageView: UIImageView!

    
    let url = URL(string: "https://felipeaccount.blob.core.windows.net/nwh/Doctor1.png")

    override func viewDidLoad() {
        super.viewDidLoad()
        videoImageView.kf.indicatorType = .activity
        videoImageView.kf.setImage(with: url)
        // Do any additional setup after loading the view.
    }
    

    /*
    // MARK: - Navigation

    // In a storyboard-based application, you will often want to do a little preparation before navigation
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        // Get the new view controller using segue.destination.
        // Pass the selected object to the new view controller.
    }
    */

}
