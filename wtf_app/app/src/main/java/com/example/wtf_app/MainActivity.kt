package com.example.wtf_app

import android.content.Intent
import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.example.wtf_app.databinding.ActivityMainBinding



class MainActivity : AppCompatActivity() {
    private lateinit var viewBinding: ActivityMainBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        viewBinding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(viewBinding.root)

        // Set up the listeners for take photo and video capture buttons
        viewBinding.btnCamera.setOnClickListener { openCamera() }
    }
    private  fun openCamera(){
        val intent = Intent(
            this,
            CameraActivity::class.java
        )
        startActivity(intent)
    }
}