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
    var SurfBreak : String = ""
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
    
    func addSpot(newSpot: Spot)->() {
        
        // Prepare URL
        let url = URL(string: "http://192.168.4.150:8080/spot")
        guard let requestUrl = url else { fatalError() }
        
        // Prepare URL Request Object
        var request = URLRequest(url: requestUrl)
        request.httpMethod = "POST"
        
        // HTTP Request Parameters which will be sent in HTTP Request Body
        request.setValue("application/json", forHTTPHeaderField: "Accept")
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        
        let jsonData = try? JSONEncoder().encode(newSpot)
        request.httpBody = jsonData
        let task = URLSession.shared.dataTask(with: request) { (data, response, error) in
            
            if let error = error {
                print("Error took place \(error)")
                return
            }
            
        }
        task.resume()
        
    }
    
    func updateSpot() {
        
        
    }
}


