//
//  CalendarViewController.swift
//  BeginnerTableView
//
//  Created by Felipe on 2019-01-27.
//  Copyright © 2019 Sean Allen. All rights reserved.
//

import UIKit

class CalendarViewController: UIViewController {

    @IBAction func check1(_ sender: UIButton) {
        sender.setBackgroundImage(#imageLiteral(resourceName: "Checkmark_brightgreen"), for: .normal)
    }
    @IBAction func check2(_ sender: UIButton) {
        sender.setBackgroundImage(#imageLiteral(resourceName: "Checkmark_brightgreen"), for: .normal)
    }
    @IBAction func check3(_ sender: UIButton) {
        sender.setBackgroundImage(#imageLiteral(resourceName: "Checkmark_brightgreen"), for: .normal)
    }
    override func viewDidLoad() {
        super.viewDidLoad()

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
