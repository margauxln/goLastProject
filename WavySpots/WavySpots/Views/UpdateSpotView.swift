//
//  UpdateSpotView.swift
//  WavySpots
//
//  Created by Alice Vedrenne on 18/05/2021.
//

import SwiftUI


struct UpdateSpotView: View {
    var spot: Spot
    @State private var titleForm: String
    @State private var surfbreakForm: String
    @State private var photoForm: String
    @State private var levelForm: Int
    @State private var addressForm: String
    
    init(spot: Spot) {
        self.spot = spot
        self.titleForm = self.spot.Title
        self.surfbreakForm = self.spot.SurfBreak
        self.photoForm = self.spot.Photo
        self.levelForm = self.spot.Level
        self.addressForm = self.spot.Address
    }
    
    @State var showingAlert = false
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
                        Api().updateSpot()
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
            .navigationBarTitle("Update Spot")
            
        }
    }
//}
}

//struct UpdateSpotView_Previews: PreviewProvider {
//    static var previews: some View {
//        UpdateSpotView(spot: Spot())
//    }
//}
