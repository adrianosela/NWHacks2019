//
//  VideoListScreen.swift
//  BeginnerTableView
//
//  Created by Sean Allen on 5/19/17.
//  Copyright © 2017 Sean Allen. All rights reserved.
//

import UIKit
import Alamofire
import QRCodeReader
import AVFoundation
import AppCenterAnalytics
import AppCenterPush
import AppCenterCrashes

class VideoListScreen: UIViewController {
    
    @IBOutlet weak var tableView: UITableView!
    
    @IBAction func signup(_ sender: Any) {
        //TODO: Network controller to post patient
        guard checkScanPermissions() else { return }
        
        readerVC.modalPresentationStyle = .formSheet
        readerVC.delegate               = self
        
        readerVC.completionBlock = { (result: QRCodeReaderResult?) in
            if let result = result {
                print("Completion with result: \(result.value.deletingPrefix("http://")) of type \(result.metadataType)")
                self.newPers(prescriptionId: result.value.deletingPrefix("http://"))
            }
        }
        
        present(readerVC, animated: true, completion: nil)
    }
    
    func newPers(prescriptionId: String) {
            let id = UserDefaults.standard.string(forKey: "id")
            let parameters: [String: Any] = [
                "prescription_id": prescriptionId,
                "patient_id": id
    
            ]
            print(parameters)
            Alamofire.request("http://ezpillzz.azurewebsites.net"+"/claim", method: .post, parameters: parameters,encoding: JSONEncoding.default).responseJSON {
                response in
                if let result = response.result.value {
                    let JSON = result as! NSDictionary
                    print(JSON["id"])
                    self.getAllPrescription(userId: id!)
                }
                print(response)
            }
    }
    
    lazy var reader: QRCodeReader = QRCodeReader()

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
    
    var videos: [Video] = []
    private let baseURL = "http://ezpillzz.azurewebsites.net"
    private let jsonDecoder = JSONDecoder()
    private let jsonEncoder = JSONEncoder()
    var tdata = [Tableview1]()

    override func viewDidLoad() {
        super.viewDidLoad()
        let id = UserDefaults.standard.string(forKey: "id")
        getAllPrescription(userId: id!)
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
                var info = [Tableview1]()
                for number in 0..<(listPrescription._Prescriptions.count) {
                    let dose = String(Int.random(in: 0 ..< 45))
                    let doctorId = listPrescription._Prescriptions[number].Doctor
                    let Remaining = listPrescription._Prescriptions[number].Remaining
                    let medicines = listPrescription._Prescriptions[number].Medicines
                    for (key, val) in medicines {
                        let name = key
                        let remaining = Remaining[name]
                        let t = Tableview1(name: name, dose: dose, doctorId: doctorId, remaining: remaining!)
                        info.append(t)
                    }
                }
                print(info)
                Alamofire.request("http://ezpillzz.azurewebsites.net"+"/patient_doctors/"+userId, method: .get, parameters: parameters).responseJSON {
                    response in
                    print(response)
                    if let json = response.data, let listDoctors = try? self.jsonDecoder.decode(ListDoctors.self, from: json) {
                        print(listDoctors, listDoctors.Doctors.count)
                        
                        for (number2) in 0..<(listDoctors.Doctors.count) {
                            let id = listDoctors.Doctors[number2].ID
                            let name = listDoctors.Doctors[number2].Name
                            for number2 in 0..<(info.count){
                                if(id == info[number2].doctorId!) { info[number2].doctor = name}
                            }
                        }
                        self.tdata = info
                        self.tableView.reloadData()
                        print("AAAA", info)
                        
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
        return tdata.count
    }
    
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let video = tdata[indexPath.row]
        let cell = tableView.dequeueReusableCell(withIdentifier: "VideoCell") as! VideoCell
        
        if(tdata.count > 0) {cell.setPrescription(data: tdata[indexPath.row])}
        
        return cell
    }
}






extension VideoListScreen: QRCodeReaderViewControllerDelegate {
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

