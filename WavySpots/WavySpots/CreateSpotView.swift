//
//  CreateSpotView.swift
//  WavySpots
//
//  Created by Fanny Armand on 11/05/2021.
//

import SwiftUI
//
//struct SurfBreak {
//    static let allSurfBreaks = [
//        "Point Break",
//        "Reef Break",
//        "River Bar",
//        "Rivermouth Break",
//        "Jetty Break",
//        "Outer Banks",
//    ]
//}
//
//struct CreateSpotView: View {
//    @State var destinationForm: String = ""
//    @State private var surfbreakForm: String = ""
//    @State var photoForm : String = ""
//    @State private var showingAlert = false
//    var body: some View {
//        NavigationView {
//            
//            Form {
//                Section {
//                    TextField("Destination", text: $destinationForm)
//                    Picker(selection: $surfbreakForm,
//                           label: Text("Surf Break")) {
//                        ForEach(SurfBreak.allSurfBreaks, id: \.self) { surfBreak in
//                            Text(surfBreak).tag(surfBreak)
//                        }
//                    }
//                    TextField("Photo", text: $photoForm)
//                }
//                Section {
//                    Button(action: {
//                        Apipost(photoForm: photoForm,surfbreakForm: [surfbreakForm], destinationForm: destinationForm).addSpot()
//                        print("Save Spot")
//                        showingAlert = true
//                        photoForm = ""
//                        surfbreakForm = ""
//                        destinationForm = ""
//                    })
//                    {
//                        HStack {
//                            Text("Save Spot")
//                        }
//                        .alert(isPresented:$showingAlert) {
//                            Alert(
//                                title: Text("Your spot has been added")
//                                //primaryButton: .destructive(Text("Delete")) {
//                                //   print("Deleting...")
//                                //}
//                            )
//                        }
//                        //.frame(width: 50, height: 100, alignment: .center)
//                        .padding(10.0)
//                        .overlay(
//                            RoundedRectangle(cornerRadius: 10.0)
//                                .stroke(lineWidth: 2.0)
//                        )}
//                    //SaveSpot.center.x = self.view.center.x
//                    
//                }
//                
//            }
//            .navigationBarTitle("Add Spot")
//            
//        }
//    }

//struct CreateSpotView_Previews: PreviewProvider {
//    static var previews: some View {
//        CreateSpotView()
//    }
//}


