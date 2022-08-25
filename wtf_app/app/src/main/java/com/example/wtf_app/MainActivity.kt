@file:Suppress("DEPRECATION")

package com.example.wtf_app

import android.Manifest
import android.R.attr
import android.app.Activity
import android.content.Intent
import android.content.pm.PackageManager
import android.graphics.Bitmap
import android.net.Uri
import android.os.Bundle
import android.os.Environment
import android.provider.MediaStore
import android.util.Log
import android.view.View
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import androidx.core.content.ContextCompat
import androidx.core.net.toUri
import com.android.volley.DefaultRetryPolicy
import com.android.volley.Response
import com.android.volley.toolbox.Volley
import com.example.wtf_app.databinding.ActivityMainBinding
import org.json.JSONArray
import org.json.JSONObject
import java.io.ByteArrayOutputStream
import java.io.File
import java.io.FileOutputStream
import java.io.IOException


const val REQUEST_CODE = 100
const val REQUEST_CAMERA = 200


class MainActivity : AppCompatActivity() {
    private lateinit var viewBinding: ActivityMainBinding
    private var imageData: ByteArray? = null
    private val postURL: String = "http://154.215.14.243:8080/post-image"

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        viewBinding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(viewBinding.root)

        viewBinding.btnCamera.setOnClickListener { capturePhoto() }
        viewBinding.btnPick.setOnClickListener { openImageGallery() }
    }

    @Suppress("DEPRECATED_IDENTITY_EQUALS")
    private fun capturePhoto() {

        val cameraIntent = Intent(MediaStore.ACTION_IMAGE_CAPTURE)
    if (ContextCompat.checkSelfPermission(this,
            Manifest.permission.CAMERA) !==
        PackageManager.PERMISSION_GRANTED) {

            ActivityCompat.requestPermissions(this,
                arrayOf(Manifest.permission.CAMERA,Manifest.permission.WRITE_EXTERNAL_STORAGE,Manifest.permission.READ_EXTERNAL_STORAGE), 1)

    } else{
        startActivityForResult(cameraIntent, REQUEST_CAMERA)
    }
    }

    private fun openImageGallery() {
        val intent = Intent(Intent.ACTION_PICK)
        intent.type = "image/*"
        startActivityForResult(intent, REQUEST_CODE)
    }
    private fun uploadImage() {
        imageData?: return
        startProgressbar()
        val request = object : VolleyFileUploadRequest(
            Method.POST,
            postURL,
            Response.Listener {res->
                val jsonArray = JSONArray(String(res.data,charset("UTF-8")))
                val datalist = mutableListOf<JSONObject>()


                for(i in 0..jsonArray.length()-1){
                    datalist.add(jsonArray.getJSONObject(i))
                }
                stopProgressbar()
                viewBinding.textResult.text = "Result: ${datalist[0].getString("class") }"
                viewBinding.textProb.text = "Probability: %.2f".format(datalist[0].getDouble("probability"))
                },
            Response.ErrorListener {
                viewBinding.textResult.text = "Error"
            }
        ) {
            override fun getByteData(): MutableMap<String, FileDataPart> {
                var params = HashMap<String, FileDataPart>()
                params["image"] = FileDataPart("image", imageData!!, "jpg")
                return params
            }
        }
        request.setRetryPolicy(
            DefaultRetryPolicy(
                12000,
                DefaultRetryPolicy.DEFAULT_MAX_RETRIES,
                DefaultRetryPolicy.DEFAULT_BACKOFF_MULT
            )
        )
        Volley.newRequestQueue(this).add(request)
    }

    @Throws(IOException::class)
    private fun createImageData(uri: Uri) {
        val inputStream = contentResolver.openInputStream(uri)
        inputStream?.buffered()?.use {
            imageData = it.readBytes()
        }
    }

    override fun onActivityResult(requestCode: Int, resultCode: Int, data: Intent?) {
        super.onActivityResult(requestCode, resultCode, data)
        if (resultCode == Activity.RESULT_OK && requestCode == REQUEST_CODE) {
            val uri = data?.data

            if (uri != null) {
                viewBinding.imgPreview.setImageURI(uri)
                createImageData(uri)
                uploadImage()
            }
        }else if (requestCode === REQUEST_CAMERA) {
            val bytes = ByteArrayOutputStream()
            val thumbnail = data?.extras?.get("data") as Bitmap
            viewBinding.imgPreview.setImageBitmap(thumbnail)
            thumbnail.compress(Bitmap.CompressFormat.JPEG, 90, bytes)
//            val destination = File(Environment.getExternalStorageDirectory(), "temp.jpg")
//            val fo: FileOutputStream
//            try {
//                fo = FileOutputStream(destination)
//                fo.write(bytes.toByteArray())
//                fo.close()
//            } catch (e: IOException) {
//                e.printStackTrace()
//            }
                imageData = bytes.toByteArray()
            uploadImage()
        }
    }
    fun stopProgressbar(){
        viewBinding.progressBar.visibility = View.INVISIBLE
        viewBinding.textResult.visibility = View.VISIBLE
        viewBinding.textProb.visibility = View.VISIBLE
    }
    fun startProgressbar(){
        viewBinding.progressBar.visibility = View.VISIBLE
        viewBinding.textResult.visibility = View.INVISIBLE
        viewBinding.textProb.visibility = View.INVISIBLE
    }
    }




