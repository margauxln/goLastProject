//
//  spotModel.swift
//  WavySpots
//
//  Created by Marine Luciani on 04/05/2021.
//

import Foundation
import SwiftUI


struct Spot: Hashable, Codable {
    var ID: String = "0"
    var Title: String = "Vide"
    var Address : String = "Vide"
    var Photo : String = ""
    var Level: Int = 0
    var SurfBreak : [String] = [""]
}


class Api {
    func getSpots(completion: @escaping([Spot]) -> ()) {
        guard let url = URL(string:"http://192.168.4.150:8080/spots") else {return}
        URLSession.shared.dataTask(with: url) { (data,_,_)in
            let decoder = JSONDecoder()
            let spots = try? decoder.decode([Spot].self, from: data!)
            DispatchQueue.main.async {
                completion(spots!)
            }
        }
        .resume()
    }
}


//class Apipost {
//    var photoForm = ""
//    var surfbreakForm = [""]
//    var destinationForm = ""
//    
//    init (photoForm: String,surfbreakForm: [String],destinationForm: String){
//        self.photoForm = photoForm
//        self.surfbreakForm = surfbreakForm
//        self.destinationForm = destinationForm}
//    
//    
//    func addSpot()->() {
//        
//        // Prepare URL
//        let url = URL(string: "http://192.168.4.150:8080/spots")
//        guard let requestUrl = url else { fatalError() }
//        
//        // Prepare URL Request Object
//        var request = URLRequest(url: requestUrl)
//        request.httpMethod = "POST"
//        
//        // HTTP Request Parameters which will be sent in HTTP Request Body
//        request.setValue("application/json", forHTTPHeaderField: "Accept")
//        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
//        
//        let newPhoto = Photos(url: self.photoForm)
//        let newField = Fields(Destination: self.destinationForm
//                              , Photos: [newPhoto], Surfbreak: self.surfbreakForm)
//        let newSpot = Spot(fields: newField)
//        let jsonData = try? JSONEncoder().encode(newSpot)
//        request.httpBody = jsonData
//        let task = URLSession.shared.dataTask(with: request) { (data, response, error) in
//            
//            if let error = error {
//                print("Error took place \(error)")
//                return
//            }
//            guard let data = data else {return}
//            do{
//                let string = String(decoding:data,as: UTF8.self)
//                print("coucou")
//                print (string)
//                let newSpot = try JSONDecoder().decode(Spot.self, from: data)
//                print(newSpot)
//                print("Response data:\n \(newSpot)")
//                print("todoItemModel Title: \(newSpot.fields)")
//                //print("todoItemModel id: \(newSpot.id)")
//            }catch let jsonErr{
//                print(jsonErr)
//            }
//            
//        }
//        task.resume()
//        
//    }
//}
