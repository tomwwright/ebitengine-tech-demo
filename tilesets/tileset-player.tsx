<?xml version="1.0" encoding="UTF-8"?>
<tileset version="1.10" tiledversion="1.11.0" name="player" tilewidth="16" tileheight="16" tilecount="12" columns="3">
 <image source="tilesheet-player.png" width="48" height="64"/>
 <tile id="0">
  <properties>
   <property name="animationName" value="idle"/>
  </properties>
  <animation>
   <frame tileid="0" duration="600"/>
   <frame tileid="1" duration="600"/>
  </animation>
 </tile>
 <tile id="3">
  <properties>
   <property name="animationName" value="walk"/>
  </properties>
  <animation>
   <frame tileid="3" duration="200"/>
   <frame tileid="4" duration="200"/>
   <frame tileid="5" duration="200"/>
   <frame tileid="4" duration="200"/>
  </animation>
 </tile>
 <tile id="6">
  <properties>
   <property name="animationName" value="walkUp"/>
  </properties>
  <animation>
   <frame tileid="7" duration="200"/>
   <frame tileid="8" duration="200"/>
   <frame tileid="9" duration="200"/>
   <frame tileid="6" duration="200"/>
  </animation>
 </tile>
</tileset>
