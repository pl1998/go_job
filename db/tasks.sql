/*
 Navicat Premium Data Transfer

 Source Server         : sqlite
 Source Server Type    : SQLite
 Source Server Version : 3030001
 Source Schema         : admin

 Target Server Type    : SQLite
 Target Server Version : 3030001
 File Encoding         : 65001

 Date: 05/08/2021 14:44:49
*/

PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for tasks
-- ----------------------------
DROP TABLE IF EXISTS "tasks";
CREATE TABLE "tasks" (
  "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  "task_name" TEXT(255),
  "status" INTEGER,
  "textarea" TEXT(255),
  "cycle" TEXT(255),
  "created_at" TIMESTAMP,
  "op_name" TEXT(255),
  "type" INTEGER,
  "email" TEXT,
  "updated_at" TIMESTAMP,
  "deleted_at" TIMESTAMP,
  "job_id" INTEGER(11)
);

-- ----------------------------
-- Auto increment value for tasks
-- ----------------------------
UPDATE "admin"."sqlite_sequence" SET seq = 4 WHERE name = 'tasks';

PRAGMA foreign_keys = true;
