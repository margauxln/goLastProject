//
//  MainView.swift
//  WavySpots
//
//  Created by Fanny Armand on 11/05/2021.
//

import SwiftUI

struct MainView: View {
    var body: some View {
        TabView {
            HomeView()
                .tabItem {
                    Label("Menu", systemImage: "list.dash")
                }
            
            CreateSpotView()
                .tabItem {
                    Label("Add Spot", systemImage: "square.and.pencil")
                }
        }
    }
}

//struct MainView_Previews: PreviewProvider {
//    static var previews: some View {
//        MainView()
//    }
//}
