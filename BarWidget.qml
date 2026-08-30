import QtQuick
import Quickshell
import qs.Commons
import qs.Ui

BarWidget {
  id: root
  moduleName: "io.github.inde3d100.homearchy"

  readonly property string pluginId: "io.github.inde3d100.homearchy"
  readonly property var homearchyService: bar?.shell?.serviceFor(pluginId)
  readonly property bool engineReady: homearchyService ? homearchyService.ready : false
  readonly property string engineState: homearchyService ? homearchyService.engineState : "unavailable"
  readonly property string engineMode: homearchyService ? homearchyService.mode : "idle"
  readonly property string statusIcon: engineState === "error" ? "󰅚" : engineReady ? "󰍽" : "󰔟"
  readonly property string statusText: engineReady
    ? (engineMode === "idle" ? "Homearchy ready" : "Homearchy · " + engineMode.replace(/_/g, " "))
    : "Homearchy · " + engineState

  property bool popupOpen: false

  function close() { popupOpen = false }
  function runMode(arguments) {
    if (homearchyService) homearchyService.runMode(arguments)
  }

  implicitWidth: button.implicitWidth
  implicitHeight: button.implicitHeight

  BarIconButton {
    id: button
    anchors.fill: parent
    bar: root.bar
    text: root.statusIcon
    slotSize: Style.bar.statusSlot
    fontSize: Style.font.icon
    tooltipText: root.statusText
    onPressed: function(buttonCode) {
      if (buttonCode === Qt.LeftButton || buttonCode === Qt.RightButton)
        root.popupOpen = !root.popupOpen
      else if (buttonCode === Qt.MiddleButton)
        root.runMode(["hints"])
    }
  }

  PopupCard {
    id: popup
    anchorItem: root
    bar: root.bar
    owner: root
    open: root.popupOpen
    contentWidth: popup.fittedContentWidth(Math.max(content.implicitWidth, Style.space(320)))
    contentHeight: popup.fittedContentHeight(content.implicitHeight)

    Column {
      id: content
      anchors.fill: parent
      spacing: Style.space(10)

      Text {
        text: "Homearchy"
        color: root.bar.foreground
        font.family: root.bar.fontFamily
        font.pixelSize: Style.font.subtitle
        font.bold: true
      }

      Text {
        width: parent.width
        text: root.statusText
        color: root.engineReady ? root.bar.foreground : Color.urgent
        font.family: root.bar.fontFamily
        font.pixelSize: Style.font.bodySmall
        elide: Text.ElideRight
      }

      PanelSeparator {
        width: parent.width
        foreground: root.bar.foreground
      }

      Grid {
        columns: 2
        spacing: Style.space(6)

        Button {
          width: Style.space(148)
          text: "Hints"
          iconText: "󰍽"
          foreground: root.bar.foreground
          enabled: root.engineReady
          leftAlign: true
          onClicked: { root.runMode(["hints"]); root.close() }
        }
        Button {
          width: Style.space(148)
          text: "Search"
          iconText: "󰍉"
          foreground: root.bar.foreground
          enabled: root.engineReady
          leftAlign: true
          onClicked: { root.runMode(["hints", "--search"]); root.close() }
        }
        Button {
          width: Style.space(148)
          text: "Grid"
          iconText: "󰕰"
          foreground: root.bar.foreground
          enabled: root.engineReady
          leftAlign: true
          onClicked: { root.runMode(["grid"]); root.close() }
        }
        Button {
          width: Style.space(148)
          text: "Recursive grid"
          iconText: "󰘕"
          foreground: root.bar.foreground
          enabled: root.engineReady
          leftAlign: true
          onClicked: { root.runMode(["recursive_grid"]); root.close() }
        }
        Button {
          width: Style.space(148)
          text: "Scroll"
          iconText: "󰹹"
          foreground: root.bar.foreground
          enabled: root.engineReady
          leftAlign: true
          onClicked: { root.runMode(["scroll"]); root.close() }
        }
        Button {
          width: Style.space(148)
          text: "Monitor"
          iconText: "󰍹"
          foreground: root.bar.foreground
          enabled: root.engineReady
          leftAlign: true
          onClicked: { root.runMode(["monitor_select"]); root.close() }
        }
      }

      PanelSeparator {
        width: parent.width
        foreground: root.bar.foreground
      }

      Button {
        width: parent.width
        text: "Restart engine"
        iconText: "󰑐"
        foreground: root.bar.foreground
        leftAlign: true
        onClicked: {
          if (root.homearchyService) root.homearchyService.requestRestart()
          root.close()
        }
      }
    }
  }
}
