//
//  SignUpViewController.swift
//  BeginnerTableView
//
//  Created by Felipe on 2019-01-27.
//  Copyright © 2019 Sean Allen. All rights reserved.
//

import UIKit
import QRCodeReader
import AVFoundation
import Alamofire
import AppCenter
import AppCenterAnalytics

import AppCenterCrashes

class SignUpViewController: UIViewController {

    @IBOutlet weak var name: UITextField!
    @IBOutlet weak var phone: UITextField!
    @IBOutlet weak var age: UITextField!
    @IBOutlet weak var email: UITextField!
    @IBOutlet weak var pass: UITextField!
    @IBOutlet weak var cpass: UITextField!
    
    @objc func dismissKeyboard(_ sender: UITapGestureRecognizer) {
        //Causes the view (or one of its embedded text fields) to resign the first responder status and drop into background
        view.endEditing(true)
    }
    
    func signupNewUser(prescriptionId: String, name: String, email: String, phone: String, age:Int) {
        let parameters: [String: Any] = [
            "prescription_id": prescriptionId,
            "name": name,
            "email": email,
            "phone": phone,
            "age": age
        ]
        print(parameters)
        Alamofire.request("http://ezpillzz.azurewebsites.net"+"/patient", method: .post, parameters: parameters,encoding: JSONEncoding.default).responseJSON {
            response in
            if let result = response.result.value {
                let JSON = result as! NSDictionary
                print(JSON["id"])
                UserDefaults.standard.set(JSON["id"]!, forKey: "id")
                self.showGame()
            }
            print(response)
        }
    }
    
    private func showGame(){
            self.performSegue(withIdentifier: "login", sender: self)
    }
    
    @IBAction func signup(_ sender: Any) {
 
        if(!((name?.text) != "") || !((phone?.text) != "") || !((age?.text) != "") || !((email?.text) != "") || !((pass?.text) != "") || !((cpass?.text) != "")){
            return
        }
        
        //TODO: Network controller to post patient
        guard checkScanPermissions() else { return }
        
        readerVC.modalPresentationStyle = .formSheet
        readerVC.delegate               = self
        
        readerVC.completionBlock = { (result: QRCodeReaderResult?) in
            if let result = result {
                print("Completion with result: \(result.value.deletingPrefix("http://")) of type \(result.metadataType)")
                self.signupNewUser(prescriptionId: result.value.deletingPrefix("http://"), name: self.name.text!, email: self.email.text!, phone: self.phone.text!, age: Int(self.age.text!)!)
            }
        }
        
        present(readerVC, animated: true, completion: nil)
  }
    
//    func compHandler(string)
    
    lazy var reader: QRCodeReader = QRCodeReader()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        let tapGesture = UITapGestureRecognizer(target: self, action: #selector(self.dismissKeyboard (_:)))
        self.view.addGestureRecognizer(tapGesture)
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
    
    lazy var readerVC: QRCodeReaderViewController = {
        let builder = QRCodeReaderViewControllerBuilder {
            $0.reader                  = QRCodeReader(metadataObjectTypes: [.qr], captureDevicePosition: .back)
            $0.showTorchButton         = true
            $0.preferredStatusBarStyle = .lightContent
            $0.showOverlayView        = true
            $0.rectOfInterest          = CGRect(x: 0.2, y: 0.2, width: 0.6, height: 0.6)
            
            $0.reader.stopScanningWhenCodeIsFound = false
        }
        
        return QRCodeReaderViewController(builder: builder)
    }()

}

extension String {
    func deletingPrefix(_ prefix: String) -> String {
        guard self.hasPrefix(prefix) else { return self }
        return String(self.dropFirst(prefix.count))
    }
}

extension SignUpViewController: QRCodeReaderViewControllerDelegate {
    func reader(_ reader: QRCodeReaderViewController, didScanResult result: QRCodeReaderResult) {
        reader.stopScanning()
        print("aasd")
        dismiss(animated: true) { [weak self] in
            let alert = UIAlertController(title: nil, message: "Please wait...", preferredStyle: .alert)
            
            let loadingIndicator = UIActivityIndicatorView(frame: CGRect(x: 10, y: 5, width: 50, height: 50))
            loadingIndicator.hidesWhenStopped = true
            loadingIndicator.style = UIActivityIndicatorView.Style.gray
            loadingIndicator.startAnimating();
            
            alert.view.addSubview(loadingIndicator)
            MSAnalytics.trackEvent("Prescription Added")
            alert.addAction(UIAlertAction(title: "OK", style: .cancel, handler: nil))
            
            self?.present(alert, animated: true, completion: nil)
        }
    }
    
    func readerDidCancel(_ reader: QRCodeReaderViewController) {
        reader.stopScanning()
        dismiss(animated: true, completion: nil)
    }
    
    func reader(_ reader: QRCodeReaderViewController, didSwitchCamera newCaptureDevice: AVCaptureDeviceInput) {
        print("Switching capture to: \(newCaptureDevice.device.localizedName)")
    }
    
    
    private func checkScanPermissions() -> Bool {
        do {
            return try QRCodeReader.supportsMetadataObjectTypes()
        } catch let error as NSError {
            let alert: UIAlertController
            
            switch error.code {
            case -11852:
                alert = UIAlertController(title: "Error", message: "This app is not authorized to use Back Camera.", preferredStyle: .alert)
                
                alert.addAction(UIAlertAction(title: "Setting", style: .default, handler: { (_) in
                    DispatchQueue.main.async {
                        if let settingsURL = URL(string: UIApplication.openSettingsURLString) {
                            UIApplication.shared.openURL(settingsURL)
                        }
                    }
                }))
                
                alert.addAction(UIAlertAction(title: "Cancel", style: .cancel, handler: nil))
            default:
                alert = UIAlertController(title: "Error", message: "Reader not supported by the current device", preferredStyle: .alert)
                alert.addAction(UIAlertAction(title: "OK", style: .cancel, handler: nil))
            }
            
            present(alert, animated: true, completion: nil)
            
            return false
        }
    }
}

