//
//  CreateSpotView.swift
//  WavySpots
//
//  Created by Fanny Armand on 11/05/2021.
//

import SwiftUI

struct SurfBreak {
    static let allSurfBreaks = [
        "Point Break",
        "Reef Break",
        "River Bar",
        "Rivermouth Break",
        "Jetty Break",
        "Outer Banks",
    ]
}

struct Level {
    static let allLevels = [1, 2, 3, 4, 5]
}

struct CreateSpotView: View {
    @State var titleForm: String = ""
    @State private var surfbreakForm: String = ""
    @State var photoForm : String = ""
    @State var levelForm: Int = 3
    @State var addressForm: String = ""
    @State private var showingAlert = false
    var body: some View {
        NavigationView {
            
            Form {
                Section {
                    TextField("Destination", text: $titleForm)
                    Picker(selection: $surfbreakForm,
                           label: Text("Surf Break")) {
                        ForEach(SurfBreak.allSurfBreaks, id: \.self) { surfBreak in
                            Text(surfBreak).tag(surfBreak)
                        }
                    }
                    TextField("Photo", text: $photoForm)
                    
                    Picker(selection: $levelForm,
                           label: Text("Difficulty Level :")) {
                        ForEach(Level.allLevels, id: \.self) { level in
                            Text(String(level))
                        }
                    }
                    TextField("Address", text: $addressForm)
                    
                }
                Section {
                    Button(action: {
                        Api().addSpot(newSpot : Spot(ID: UUID().uuidString, Title: titleForm, Address: addressForm, Photo: photoForm, Level: levelForm, SurfBreak: surfbreakForm))
                        print("Save Spot")
                        showingAlert = true
                    })
                    {
                        HStack {
                            Text("Save Spot")
                        }
                        .alert(isPresented:$showingAlert) {
                            Alert(
                                title: Text("Your spot has been added")
                                //primaryButton: .destructive(Text("Delete")) {
                                //   print("Deleting...")
                                //}
                            )
                        }
                        //.frame(width: 50, height: 100, alignment: .center)
                        .padding(10.0)
                        .overlay(
                            RoundedRectangle(cornerRadius: 10.0)
                                .stroke(lineWidth: 2.0)
                        )}
                    //SaveSpot.center.x = self.view.center.x
                    
                }
                
            }
            .navigationBarTitle("Add Spot")
            
        }
    }
}

//struct CreateSpotView_Previews: PreviewProvider {
//    static var previews: some View {
//        CreateSpotView()
//    }
//}


