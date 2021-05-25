//
//  DestinationPageView.swift
//  WavySpots
//
//  Created by Marine Luciani on 04/05/2021.
//

import SwiftUI

struct DestinationPageView: View {
    var spot : Spot
    @State var isModal: Bool = false
    
    init(spot: Spot){self.spot = spot}
    
    var body: some View {
        Text(spot.Address)
            .font(.title)
            .foregroundColor(Color("Darkblue"))
        Text("Difficulty : " + String(spot.Level))
        Text(spot.SurfBreak)
            .font(.subheadline)
        CircleImage(photo:spot.Photo)
        MapView()
        
        Button("Update :(") {
            self.isModal = true
        }.sheet(isPresented: $isModal, content: {
            UpdateSpotView(spot:spot)
        })
    }
}




//struct DestinationPageView_Previews: PreviewProvider {
//    static var previews: some View {
//        DestinationPageView(id:1)
//    }
//}
