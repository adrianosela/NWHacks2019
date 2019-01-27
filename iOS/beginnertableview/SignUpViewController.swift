//
//  SignUpViewController.swift
//  BeginnerTableView
//
//  Created by Felipe on 2019-01-27.
//  Copyright © 2019 Sean Allen. All rights reserved.
//

import UIKit

class SignUpViewController: UIViewController {

    @IBOutlet weak var name: UITextField!
    @IBOutlet weak var phone: UITextField!
    @IBOutlet weak var age: UITextField!
    @IBOutlet weak var email: UITextField!
    @IBOutlet weak var pass: UITextField!
    @IBOutlet weak var cpass: UITextField!
    
    @IBAction func signup(_ sender: Any) {
        if(!((name?.text) != nil) || !((phone?.text) != nil) || !((age?.text) != nil) || !((email?.text) != nil) || !((pass?.text) != nil) || !((cpass?.text) != nil)){
            return
        }
        //TODO: Network controller to post patient
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
